package main

import (
	"context"
	"errors"
	"net"

	"github.com/bmg02/orchestrator-service/service"
	"google.golang.org/grpc"
)

type User struct {
	name  string
	class string
	roll  int64
}

func main() {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	s := service.User{}

	service.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}

func (s *User) GetUserByName(ctx context.Context, req *service.UserRequest) (*service.UserResponse, error) {
	return nil, errors.New("Hey there")
}
