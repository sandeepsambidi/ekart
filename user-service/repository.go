package main

import (
	"context"
	"log"

	pb "github.com/sandeepsambidi/ekart/user-service/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository has actions for a user
type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

//UserRepository is impl of Repository in User
type UserRepository struct {
	userCollection *mongo.Collection
}

//GetAll returns all users
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	cursor, err := repo.userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Panicf("Failed to get all users cursor: %v", err)
	}

	for cursor.Next(context.Background()) {
		var user *pb.User

		decodeError := cursor.Decode(&user)
		if decodeError != nil {
			return nil, decodeError
		}
		users = append(users, user)
	}

	return users, err

}

//Get returns a user by id
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	userresult := repo.userCollection.FindOne(context.Background(), bson.M{"id": id})
	err := userresult.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//Create will insert a new user document in user collection
func (repo *UserRepository) Create(user *pb.User) error {
	_, err := repo.userCollection.InsertOne(context.Background(), user)
	return err
}

//GetByEmailAndPassword is used for auth by validating email and password
func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	//TODO
	return nil, nil
}
