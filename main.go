package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	pb "test/klaus/proto"

	"github.com/fsnotify/fsnotify"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var (
	port             = flag.Int("port", 9999, "The server port")
	databaseInMemory *sql.DB //should not be global, should be sth that can be passed to grpc server functions, but good enough for the example
)

type server struct {
	pb.UnimplementedAggregateScoresServer
	pb.UnimplementedChangeInScoreServer
	pb.UnimplementedOverallScoreServer
	pb.UnimplementedTicketScoresServer
}

type ratingCategory struct {
	ticket_id  int32
	name       string
	rating     int
	weight     float64
	created_at string
}

func loadDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db?cache=shared&mode=memory")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getWeightedScore(rating int, weight float64) float64 {
	score := float64(rating) * 100 / 5
	return score * weight
}

func getData(start string, end string, orderBy string) []ratingCategory {
	row, err := databaseInMemory.Query(fmt.Sprintf(`select ratings.ticket_id, ratings.rating, rating_categories.name,
	rating_categories.weight, ratings.created_at from ratings 
	left join rating_categories 
	on ratings.rating_category_id = rating_categories.id
	where ratings.created_at >= '%s' and ratings.created_at <= '%sT23:59:59' ORDER by %s`, start, end, orderBy))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var result []ratingCategory

	for row.Next() {
		var rating ratingCategory
		err := row.Scan(&rating.ticket_id, &rating.rating, &rating.name, &rating.weight, &rating.created_at)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, rating)

	}
	return result
}

func (s *server) SendAggregateScores(ctx context.Context, in *pb.AggregateScoresRequest) (*pb.AggregateScoresReply, error) {
	result := getData(in.GetDateStart(), in.GetDateEnd(), "rating_categories.name")

	var prevName = ""
	var count = 0

	var categories []*pb.AggregateScoresCategory
	var dates []*pb.AggregateScoresCategoriesDate
	var cat = new(pb.AggregateScoresCategory)

	for _, element := range result {
		date := new(pb.AggregateScoresCategoriesDate)
		date.Date = element.created_at
		date.Percentage = getWeightedScore(element.rating, element.weight)
		dates = append(dates, date)
		if element.name != prevName {
			if count != 0 {
				cat.Ratings = int32(count)
				cat.Dates = dates
				dates = nil
			}
			cat = new(pb.AggregateScoresCategory)
			cat.Category = element.name
			categories = append(categories, cat)
			count = 0
		}
		prevName = element.name
		count++
	}
	cat.Ratings = int32(count)
	cat.Dates = dates

	return &pb.AggregateScoresReply{Categories: categories}, nil
}

func (s *server) SendTicketScores(ctx context.Context, in *pb.TicketScoresRequest) (*pb.TicketScoresReply, error) {
	result := getData(in.GetDateStart(), in.GetDateEnd(), "ratings.ticket_id")

	var tickets []*pb.TicketScoresItem

	for _, element := range result {
		ticket := new(pb.TicketScoresItem)
		ticket.Id = element.ticket_id
		var ticketCategories []*pb.TicketScoresCategory
		for _, ticketWithId := range result {
			if ticketWithId.ticket_id == ticket.Id {
				ticketCategory := new(pb.TicketScoresCategory)
				ticketCategory.Name = ticketWithId.name
				var scores []float64
				for _, ticketWithCat := range result {
					if ticketWithCat.name == ticketCategory.Name {
						percentage := getWeightedScore(element.rating, element.weight)
						scores = append(scores, percentage)
					}
				}
				total := 0.0
				for _, score := range scores {
					total = total + score
				}
				average := total / float64(len(scores))
				ticketCategory.Percentage = average
				ticketCategories = append(ticketCategories, ticketCategory)
			}
		}
		ticket.Categories = append(ticket.Categories, ticketCategories...)
		tickets = append(tickets, ticket)
	}

	return &pb.TicketScoresReply{Tickets: tickets}, nil
}

func (s *server) SendOverallScore(ctx context.Context, in *pb.OverallScoreRequest) (*pb.OverallScoreReply, error) {
	result := getData(in.GetDateStart(), in.GetDateEnd(), "ratings.ticket_id")
	var scores []float64

	for _, element := range result {
		percentage := getWeightedScore(element.rating, element.weight)
		scores = append(scores, percentage)
	}
	total := 0.0
	for _, score := range scores {
		total = total + score
	}
	average := total / float64(len(scores))

	return &pb.OverallScoreReply{Score: average}, nil
}

func (s *server) SendChangeInScore(ctx context.Context, in *pb.ChangeInScoreRequest) (*pb.ChangeInScoreReply, error) {
	firstPeriodData := getData(in.GetFromDateStart(), in.GetFromDateEnd(), "ratings.ticket_id")
	secondPeriodData := getData(in.GetToDateStart(), in.GetToDateEnd(), "ratings.ticket_id")

	var scoresFirst []float64

	for _, element := range firstPeriodData {
		percentage := getWeightedScore(element.rating, element.weight)
		scoresFirst = append(scoresFirst, percentage)
	}
	totalFirst := 0.0
	for _, score := range scoresFirst {
		totalFirst = totalFirst + score
	}
	averageFirst := totalFirst / float64(len(scoresFirst))

	var scoresSecond []float64

	for _, element := range secondPeriodData {
		percentage := getWeightedScore(element.rating, element.weight)
		scoresSecond = append(scoresSecond, percentage)
	}
	total := 0.0
	for _, score := range scoresFirst {
		total = total + score
	}
	averageSecond := total / float64(len(scoresSecond))

	difference := ((averageFirst - averageSecond) / (averageFirst + averageSecond) / 2) * 100

	return &pb.ChangeInScoreReply{Change: difference}, nil
}

func main() {
	fmt.Println("Server started")

	fmt.Println("Loading database to memory")
	databaseInMemory = loadDatabase()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("Database modified, reloading data to memory")
					databaseInMemory = loadDatabase()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./database.db")
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAggregateScoresServer(s, &server{})
	pb.RegisterTicketScoresServer(s, &server{})
	pb.RegisterOverallScoreServer(s, &server{})
	pb.RegisterChangeInScoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	<-make(chan struct{})
}
