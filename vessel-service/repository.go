package main

import (
	"context"
	"errors"

	pb "github.com/sandeepsambidi/ekart/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository is an interface for Vessel collection
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
}

//VesselRepository is an impl of Repository
type VesselRepository struct {
	vesselCollection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	cursor, err := repo.vesselCollection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var vessel *pb.Vessel
		err := cursor.Decode(&vessel)
		if err != nil {
			return nil, err
		}
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

//Create a new vessel in db
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := repo.vesselCollection.InsertOne(context.Background(), vessel)
	return err
}
