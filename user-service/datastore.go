package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CreateClient creates a mongo client for ekart db
func CreateClient(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
}
