package service

import (
	"context"
	"github.com/kirpepa/spendly-api/auth/model"
	"github.com/kirpepa/spendly-api/auth/proto"
	"github.com/kirpepa/spendly-api/auth/repository"
	"github.com/kirpepa/spendly-api/auth/token"
	"golang.org/x/crypto/bcrypt"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	repo *repository.UserRepo
	jwt  *token.JWTManager
}

func NewAuthServer(repo *repository.UserRepo, jwt *token.JWTManager) *AuthServer {
	return &AuthServer{repo: repo, jwt: jwt}
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	if err := s.repo.Create(user); err != nil {
		return &proto.AuthResponse{Error: "email already registered"}, nil
	}

	token, _ := s.jwt.Generate(user.ID, user.Email)
	return &proto.AuthResponse{Token: token}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return &proto.AuthResponse{Error: "invalid credentials"}, nil
	}
	token, _ := s.jwt.Generate(user.ID, user.Email)
	return &proto.AuthResponse{Token: token}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, err := s.jwt.Verify(req.Token)
	if err != nil {
		return &proto.ValidateResponse{Valid: false}, nil
	}
	return &proto.ValidateResponse{
		Valid:  true,
		UserId: uint64(claims.UserID),
		Email:  claims.Email,
	}, nil
}
