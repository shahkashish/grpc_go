package main

import (
	"log"
	"net"

	pb "github.com/shahkashish/grpc_go/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UserManagementServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
