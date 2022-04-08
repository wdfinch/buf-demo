package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go.buf.build/grpc/go/petland/dogapis/petland/dog/v1"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
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

func main() {
	file, _ := ioutil.ReadFile("database.json")
	database := Database{}
	_ = json.Unmarshal(file, &database)

	//fmt.Printf("%+v\n", database)
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	log.Println(dogv1.Breed_BREED_BULLDOG)
	//dServer := &calculatorServer{}
	//dog.RegisterDogServer(s, dServer)

	fmt.Println("starting server on port", *port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}
