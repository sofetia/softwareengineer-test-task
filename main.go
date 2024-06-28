package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	pb "test/klaus/proto"
	"time"

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
	ticket_id int32
	name      string
	score     float64
	date      string
}

func loadDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db?cache=shared&mode=memory")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getData(start string, end string) []ratingCategory {
	q := "select ratings.ticket_id, rating_categories.name, " +
		"STRFTIME('%Y-%m-%d', ratings.created_at) as date, round(ratings.rating * rating_categories.weight * 10) as score from ratings " +
		"left join rating_categories " +
		"on ratings.rating_category_id = rating_categories.id " +
		fmt.Sprintf("where ratings.created_at >= '%s' and ratings.created_at <= '%s' order by ratings.created_at", start, end)
	row, err := databaseInMemory.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var result []ratingCategory

	for row.Next() {
		var rating ratingCategory
		err := row.Scan(&rating.ticket_id, &rating.name, &rating.date, &rating.score)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, rating)

	}
	return result
}

func getStartOfWeeks(start time.Time, end time.Time) []time.Time {
	var dates []time.Time
	dates = append(dates, start)
	for start.Before(end) {
		start = start.AddDate(0, 0, 7)
		if !start.Before(end) {
			break
		}
		dates = append(dates, start)
	}
	return dates
}

func (s *server) SendAggregateScores(ctx context.Context, in *pb.AggregateScoresRequest) (*pb.AggregateScoresReply, error) {
	result := getData(in.GetDateStart().AsTime().String(), in.GetDateEnd().AsTime().String())

	isMonth := false

	if in.GetDateEnd().AsTime().Sub(in.GetDateStart().AsTime()) > 30*24*time.Hour {
		isMonth = true
	}

	var categories []*pb.AggregateScoresCategory

	for _, row := range result {
		excists := false
		for _, cat := range categories {
			if cat.Category == row.name {
				excists = true
				break
			}
		}
		if !excists {
			category := new(pb.AggregateScoresCategory)
			category.Category = row.name
			categories = append(categories, category)
		}
	}

	for index, cat := range categories {
		var dates []*pb.AggregateScoresCategoriesDate
		var ratings = 0
		if !isMonth {
			for _, row := range result {
				if row.name == cat.Category && !isMonth {
					ratings++
					excists := false
					for _, date := range dates {
						if date.Date == row.date {
							excists = true
							break
						}
					}
					if !excists {
						date := new(pb.AggregateScoresCategoriesDate)
						date.Date = row.date
						dates = append(dates, date)
					}
				}
			}
		}
		if isMonth {
			startOfWeeks := getStartOfWeeks(in.GetDateStart().AsTime(), in.GetDateEnd().AsTime())
			for _, weekday := range startOfWeeks {
				date := new(pb.AggregateScoresCategoriesDate)
				date.Date = weekday.Format("2006-01-02")
				dates = append(dates, date)
			}
		}
		categories[index].Dates = dates
		categories[index].Ratings = int32(ratings)
	}

	for index, cat := range categories {
		for index2, date := range cat.Dates {
			var scores []float64
			if !isMonth {
				for _, row := range result {
					if row.name == cat.Category && row.date == date.Date {
						scores = append(scores, row.score)
					}
				}

			}
			if isMonth {
				for _, row := range result {
					catDate, _ := time.Parse("2006-01-02", date.Date)
					rowDate, _ := time.Parse("2006-01-02", row.date)
					if row.name == cat.Category && rowDate.After(catDate) && rowDate.Before(catDate.AddDate(0, 0, 7)) {
						scores = append(scores, row.score)
					}
				}
			}
			total := 0.0
			for _, score := range scores {
				total = total + score
			}
			average := total / float64(len(scores))
			categories[index].Dates[index2].Percentage = average
		}
	}
	for index, cat := range categories {
		var scores []float64
		for _, date := range cat.Dates {
			scores = append(scores, date.Percentage)
		}
		total := 0.0
		for _, score := range scores {
			total = total + score
		}
		average := total / float64(len(scores))
		categories[index].Score = average
	}

	return &pb.AggregateScoresReply{Categories: categories}, nil
}

func (s *server) SendTicketScores(ctx context.Context, in *pb.TicketScoresRequest) (*pb.TicketScoresReply, error) {
	result := getData(in.GetDateStart().AsTime().String(), in.GetDateEnd().AsTime().String())

	var tickets []*pb.TicketScoresItem

	for _, row := range result {
		excists := false
		for _, ticket := range tickets {
			if ticket.Id == row.ticket_id {
				excists = true
				break
			}
		}
		if !excists {
			ticket := new(pb.TicketScoresItem)
			ticket.Id = row.ticket_id
			tickets = append(tickets, ticket)
		}
	}

	for index, ticket := range tickets {
		var categories []*pb.TicketScoresCategory
		for _, row := range result {
			if row.ticket_id == ticket.Id {
				excists := false
				for _, cat := range categories {
					if cat.Name == row.name {
						excists = true
						break
					}
				}
				if !excists {
					cat := new(pb.TicketScoresCategory)
					cat.Name = row.name
					categories = append(categories, cat)
				}
			}
		}
		tickets[index].Categories = categories

		for index, ticket := range tickets {
			for index2, cat := range ticket.Categories {
				var scores []float64
				for _, row := range result {
					if ticket.Id == row.ticket_id && cat.Name == row.name {
						scores = append(scores, row.score)
					}
				}
				total := 0.0
				for _, score := range scores {
					total = total + score
				}
				average := total / float64(len(scores))
				tickets[index].Categories[index2].Percentage = average
			}

		}
	}

	return &pb.TicketScoresReply{Tickets: tickets}, nil
}

func (s *server) SendOverallScore(ctx context.Context, in *pb.OverallScoreRequest) (*pb.OverallScoreReply, error) {
	result := getData(in.GetDateStart().AsTime().String(), in.GetDateEnd().AsTime().String())
	total := 0.0
	for _, element := range result {
		total = total + element.score
	}
	average := total / float64(len(result))

	return &pb.OverallScoreReply{Score: average}, nil
}

func (s *server) SendChangeInScore(ctx context.Context, in *pb.ChangeInScoreRequest) (*pb.ChangeInScoreReply, error) {
	firstPeriodData := getData(in.GetFromDateStart().AsTime().String(), in.GetFromDateEnd().AsTime().String())
	secondPeriodData := getData(in.GetToDateStart().AsTime().String(), in.GetToDateEnd().AsTime().String())

	totalFirst := 0.0

	for _, element := range firstPeriodData {
		totalFirst = totalFirst + element.score
	}
	averageFirst := totalFirst / float64(len(firstPeriodData))

	total := 0.0
	for _, element := range secondPeriodData {
		total = total + element.score
	}
	averageSecond := total / float64(len(secondPeriodData))

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
