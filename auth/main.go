package main

import (
	"github.com/kirpepa/spendly-api/auth/config"
	"github.com/kirpepa/spendly-api/auth/db"
	"github.com/kirpepa/spendly-api/auth/proto"
	"github.com/kirpepa/spendly-api/auth/repository"
	"github.com/kirpepa/spendly-api/auth/service"
	"github.com/kirpepa/spendly-api/auth/token"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	database := db.InitDB(cfg.DBUrl)
	repo := repository.NewUserRepo(database)
	jwt := token.NewJWTManager(cfg.JWTSecret, cfg.JWTExpire)

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, service.NewAuthServer(repo, jwt))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Auth service listening on :50051")
	grpcServer.Serve(lis)
}
