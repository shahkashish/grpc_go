package main

import (
	"context"

	pb "github.com/shahkashish/grpc_go/proto"
)

func (s *Server) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	println("HEREEEEEEEE")
	return &pb.User{
		Name: in.GetName(), Age: in.GetAge(), Id: 100,
	}, nil
}
