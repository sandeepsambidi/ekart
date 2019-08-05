package main

import (
	"context"
	"log"

	pbcon "github.com/sandeepsambidi/ekart/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(consignment *pbcon.Consignment) error
	GetAll() ([]*pbcon.Consignment, error)
}

// Repository - holds CRUD functions for consignment collection
type Repository struct {
	collection *mongo.Collection
}

// Create a new consignment
func (repo *Repository) Create(consignment *pbcon.Consignment) error {
	_, err := repo.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll consignments
func (repo *Repository) GetAll() ([]*pbcon.Consignment, error) {
	consCursor, err := repo.collection.Find(context.Background(), nil, nil)
	if err != nil {
		log.Fatalf("error getting cursor of consigment collection %v", err)
	}
	var consignments []*pbcon.Consignment
	for consCursor.Next(context.Background()) {
		var consignment *pbcon.Consignment
		if err := consCursor.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}

	return consignments, err
}
