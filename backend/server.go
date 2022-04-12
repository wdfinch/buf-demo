package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	dogv1 "go.buf.build/grpc/go/petland/dogapis/petland/dog/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = 50051
)

type Dog struct {
	Id                     uint32         `json:"id"`
	Name                   string         `json:"name"`
	Age                    uint32         `json:"age"`
	Breed                  string         `json:"breed"`
	BirthDate              int            `json:"birthDate"`
	DogPrice               int            `json:"dogPrice"`
	FavoriteToys           []string       `json:"favoriteToys"`
	ProfileImageUrl        string         `json:"profileImageUrl"`
	IsVaccinated           bool           `json:"isVaccinated"`
	Weight                 float32        `json:"weight"`
	PostalAddress          string         `json:"postalAddress"`
	FavoriteFoodsWithPrice map[string]int `json:"favoriteFoodsWithPrice"`
}

type Database struct {
	Dogs []Dog `json:"dogs"`
}

type Server struct {
	dogv1.UnimplementedDogApiServiceServer
}

func getDatabase() Database {
	file, _ := ioutil.ReadFile("database.json")
	database := Database{}
	_ = json.Unmarshal(file, &database)
	return database
}

func (cs *Server) GetDog(ctx context.Context, req *dogv1.GetDogRequest) (*dogv1.GetDogResponse, error) {
	fmt.Println("processing qequest")
	database := getDatabase()
	var dog *Dog = nil
	for _, d := range database.Dogs {
		if d.Id == req.DogId {
			dog = &d
		}
	}

	if dog == nil {
		err := status.Error(codes.NotFound, "No dog was found with the supplied ID")
		return nil, err
	}

	cRes := getResp(dog)

	return &cRes, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%d", ":", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dogv1.RegisterDogApiServiceServer(s, &Server{})

	fmt.Println("starting server on port", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}
