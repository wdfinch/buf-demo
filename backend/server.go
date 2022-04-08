package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.buf.build/grpc/go/petland/dogapis/petland/dog/v1"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)

var (
	port = 50052
)

type Database struct {
	Id                     int            `json:"id"`
	Name                   string         `json:"name"`
	Age                    int            `json:"age"`
	Breed                  string         `json:"breed"`
	BirthDate              int            `json:"birthDate"`
	DogPrice               int            `json:"dogPrice"`
	FavoriteToys           []string       `json:"favoriteToys"`
	ProfileImageUrl        string         `json:"profileImageUrl"`
	IsVaccinated           bool           `json:"isVaccinated"`
	Weight                 int            `json:"weight"`
	PostalAddress          string         `json:"postalAddress"`
	FavoriteFoodsWithPrice map[string]int `json:"favoriteFoodsWithPrice"`
}

type Server struct {
	dogv1.UnimplementedDogApiServiceServer
}

func (cs *Server) GetDog(ctx context.Context, req *dogv1.GetDogRequest) (*dogv1.GetDogResponse, error) {
	cRes := &dogv1.GetDogResponse{}
	fmt.Println("hello")
	return cRes, nil
}

func main() {
	file, _ := ioutil.ReadFile("database.json")
	database := Database{}
	_ = json.Unmarshal(file, &database)

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
