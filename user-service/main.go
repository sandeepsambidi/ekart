package main

import (
	"fmt"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/sandeepsambidi/ekart/user-service/proto/user"
)

func main() {
	//get the service handle
	srv := micro.NewService(
		micro.Name("user"),
	)
	srv.Init()
	uri := "mongodb"
	usercollection, err := CreateClient(uri)

	if err != nil {
		log.Panicf("Error getting db client: %v", err)
	}

	userCollection := usercollection.Database("ekart").Collection("user")
	userRepo := &UserRepository{userCollection}
	errorinreg := pb.RegisterUserServiceHandler(srv.Server(), &service{userRepo})

	if errorinreg != nil {
		log.Printf("error in registering handler: %v", errorinreg)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
