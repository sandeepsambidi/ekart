package main

import (
	"context"
	"log"

	pb "github.com/sandeepsambidi/ekart/user-service/proto/user"
)

type service struct {
	repo Repository
}

func (s *service) Create(ctx context.Context, in *pb.User, out *pb.Response) error {
	err := s.repo.Create(in)
	if err != nil {
		log.Panic("Create user failed: %v", err)
		return err
	}
	out.User = in
	return nil
}

func (s *service) Get(ctx context.Context, in *pb.User, out *pb.Response) error {
	user, err := s.repo.Get(in.GetId())
	if err != nil {
		return err
	}

	out.User = user
	return nil
}

func (s *service) GetAll(ctx context.Context, in *pb.Request, out *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	out.Users = users
	return nil
}

func (s *service) Auth(ctx context.Context, in *pb.User, out *pb.Token) error {
	_, err := s.repo.GetByEmailAndPassword(in)
	if err != nil {
		return err
	}
	out.Token = "testingabc"
	return nil
}

func (s *service) ValidateToken(ctx context.Context, in *pb.Token, out *pb.Token) error {
	//TODO
	return nil
}
