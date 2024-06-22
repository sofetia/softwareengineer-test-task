package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	pb "test/klaus/proto"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedAggregateScoresServer
	pb.UnimplementedChangeInScoreServer
	pb.UnimplementedOverallScoreServer
	pb.UnimplementedTicketScoresServer
}

type ratingCategory struct {
	name       string
	rating     int32
	weight     float32
	created_at string
}

func connectToSQLite() *sql.DB {
	sqliteDatabase, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer sqliteDatabase.Close() // Defer Closing the database
	return sqliteDatabase
}

func getWeightedScore(rating int, weight float64) float64 {
	score := float64(rating) * 100 / 5
	return score * weight
}

func getData(start string, end string) []ratingCategory {
	db := connectToSQLite()

	row, err := db.Query(fmt.Sprintf(`select ratings.rating, rating_categories.name,
	rating_categories.weight, ratings.created_at from ratings 
	left join rating_categories 
	on ratings.rating_category_id = rating_categories.id
	where ratings.created_at >= '%s' and ratings.created_at <= '%sT23:59:59' ORDER by rating_categories.name`, start, end))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var result []ratingCategory

	for row.Next() {
		var rating ratingCategory
		err := row.Scan(&rating.rating, &rating.name, &rating.weight, &rating.created_at)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, rating)

	}
	db.Close()
	return result
}

func (s *server) SendAggregateScores(ctx context.Context, in *pb.AggregateScoresRequest) (*pb.AggregateScoresReply, error) {
	result := getData(in.GetDateStart(), in.GetDateEnd())

	var prevName = ""
	var count = 0

	var categories []*pb.AggregateScoresCategory
	var dates []*pb.AggregateScoresCategoriesDate
	var cat = new(pb.AggregateScoresCategory)

	for _, element := range result {
		date := new(pb.AggregateScoresCategoriesDate)
		date.Date = element.created_at
		date.Percentage = int32(getWeightedScore(int(element.rating), float64(element.weight)))
		dates = append(dates, date)
		if element.name != prevName {
			if count != 0 {
				cat.Ratings = int32(count)
				fmt.Println(dates)
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

func main() {
	fmt.Println("Server started")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAggregateScoresServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
