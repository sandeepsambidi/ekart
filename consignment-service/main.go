package main

import (
	"fmt"

	// Import the generated protobuf code

	micro "github.com/micro/go-micro"

	pb "github.com/sandeepsambidi/ekart/consignment-service/proto/consignment"
	pbvessel "github.com/sandeepsambidi/ekart/vessel-service/proto/vessel"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

func main() {

	repo := &Repository{}

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
