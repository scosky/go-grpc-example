package services

import (
	context "context"
)

type UserService struct {
}

func (userService *UserService) GetUserInfo(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	return &UserResponse{UserId: request.UserId, Name: "詹丹的", Age: 23}, nil
}
