package main

import (
	"github.com/kirpepa/spendly-api/user/config"
	"github.com/kirpepa/spendly-api/user/db"
	"github.com/kirpepa/spendly-api/user/proto"
	"github.com/kirpepa/spendly-api/user/repository"
	"github.com/kirpepa/spendly-api/user/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	database := db.InitDB(cfg.DBUrl)
	repo := repository.NewUserRepo(database)
	userServer := service.NewUserServer(repo)

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userServer)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("User service listening on :50052")
	grpcServer.Serve(lis)
}
