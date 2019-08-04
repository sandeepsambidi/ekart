package main

import (
	"context"

	pb "github.com/sandeepsambidi/ekart/vessel-service/proto/vessel"
)

// Our grpc service handler
type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, in *pb.Vessel, out *pb.Response) error {
	err := s.repo.Create(in)
	if err != nil {
		return err
	}
	out.Created = true
	return nil
}
