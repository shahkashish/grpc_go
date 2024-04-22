package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shahkashish/grpc_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:8080"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["kashish"] = 10
	new_users["rangu"] = 5
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: "kashish", Age: 12})
		if err != nil {
			log.Fatalf("could not create user %s %d %v", name, age, err)
		}
		log.Printf(`UserDetails:
		Name : %s
		Age : %d
		ID : %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
