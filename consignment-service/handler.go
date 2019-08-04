package main

import (
	"context"
	"log"

	pbcon "github.com/sandeepsambidi/ekart/consignment-service/proto/consignment"
	pbves "github.com/sandeepsambidi/ekart/vessel-service/proto/vessel"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo         repository
	vesselClient pbves.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(con context.Context, req *pbcon.Consignment, res *pbcon.Response) error {
	//1. find a vessel for the consignment
	//2. attach vessel to the consignment
	//3. store the consignment in the DB
	specification := &pbves.Specification{
		MaxWeight: req.GetWeight(),
		Capacity:  int32(len(req.Containers))}
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), specification)

	if err != nil {
		return err
	}

	log.Printf("Found vessel: %s", vesselResponse.Vessel.GetName())

	req.VesselId = vesselResponse.Vessel.GetId()

	//create the consignment ie add a document in consignment collection
	erroncreate := s.repo.Create(req)

	if erroncreate != nil {
		return err
	}

	res.Created = true
	//res.Consignment = createdcon
	return nil
}

//get all consignments
func (s *service) GetConsignments(con context.Context, req *pbcon.GetRequest, res *pbcon.Response) error {
	consignments, err := s.repo.GetAll()

	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
