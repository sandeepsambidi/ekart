// vessel-service/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/sandeepsambidi/ekart/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

const (
	port          = ":50051"
	defaultDBHost = "datastore:27017"
)

func createDummyData(repo Repository) {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBHost
	}
	dbclientconn, dberror := CreateClient(uri)
	if dberror != nil {
		log.Panic(dberror)
	}
	defer dbclientconn.Disconnect(context.TODO())

	vescollection := dbclientconn.Database("ekart").Collection("vessel")
	log.Printf("retrieved vessel name: %s", vescollection.Name())
	repo := &VesselRepository{vescollection}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("vessel"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
