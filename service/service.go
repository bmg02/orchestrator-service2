package service

import (
	"context"
	"errors"
	// "github.com/bmg02/orchestrator-service/service"
)

type User struct {
	name  string
	class string
	roll  int64
}

func (s *User) GetUserByName(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return &UserResponse{}, errors.New("Not implemented yet. Bhuvan will implement me")
}
