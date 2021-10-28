package main

import (
	"context"
	"log"

	"github.com/bmg02/orchestrator-service/service"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := service.NewUserServiceClient(conn)

	// message :=

	response, err := c.GetUserByName(context.Background(), &service.UserRequest{})

	if err != nil {
		panic(err)
	}

	log.Printf("Response: %s", response.String())
}
