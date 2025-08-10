package handler

import (
	"context"
	"github.com/kirpepa/spendly-api/user/proto"
	"github.com/kirpepa/spendly-api/user/service"
)

type GRPCHandler struct {
	proto.UnimplementedUserServiceServer
	service *service.UserServer
}

func NewGRPCHandler(s *service.UserServer) *GRPCHandler {
	return &GRPCHandler{service: s}
}

func (h *GRPCHandler) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	return h.service.GetUser(ctx, req)
}

func (h *GRPCHandler) ListUsers(ctx context.Context, req *proto.Empty) (*proto.UserListResponse, error) {
	return h.service.ListUsers(ctx, req)
}

func (h *GRPCHandler) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	return h.service.DeleteUser(ctx, req)
}
