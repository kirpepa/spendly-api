package service

import (
	"context"
	"github.com/kirpepa/spendly-api/user/proto"
	"github.com/kirpepa/spendly-api/user/repository"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	repo *repository.UserRepo
}

func NewUserServer(repo *repository.UserRepo) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	user, err := s.repo.GetByID(uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{
		Id:    uint64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, _ *proto.Empty) (*proto.UserListResponse, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	var protoUsers []*proto.UserResponse
	for _, u := range users {
		protoUsers = append(protoUsers, &proto.UserResponse{
			Id:    uint64(u.ID),
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return &proto.UserListResponse{Users: protoUsers}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := s.repo.Delete(uint(req.UserId))
	return &proto.DeleteUserResponse{Success: err == nil}, err
}
