package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code

	micro "github.com/micro/go-micro"

	pb "github.com/sandeepsambidi/ekart/consignment-service/proto/consignment"
	pbvessel "github.com/sandeepsambidi/ekart/vessel-service/proto/vessel"
)

const (
	port          = ":50051"
	defaultDBHost = "datastore:27017"
)

func main() {

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBHost
	}
	dbclientconn, dberr := CreateClient(uri)
	if dberr != nil {
		log.Panic(dberr)
	}
	defer dbclientconn.Disconnect(context.TODO())
	consignmentCollection := dbclientconn.Database("ekart").Collection("consignments")
	repo := &Repository{consignmentCollection}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("consignment"),
	)

	// Init will parse the command line flags.
	srv.Init()
	vesselClient := pbvessel.NewVesselService("vessel", srv.Client())
	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
