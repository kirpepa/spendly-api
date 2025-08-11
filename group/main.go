package main

import (
	"log"
	"net"

	"github.com/kirpepa/spendly-api/group/proto"
	"github.com/kirpepa/spendly-api/group/repository"
	"github.com/kirpepa/spendly-api/group/service"

	"github.com/kirpepa/spendly-api/internal/config"
	"github.com/kirpepa/spendly-api/internal/db"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	database := db.InitDB(cfg.DBUrl)
	repo := repository.NewGroupRepo(database)
	groupServer := service.NewGroupServer(repo)

	grpcServer := grpc.NewServer()
	proto.RegisterGroupServiceServer(grpcServer, groupServer)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Group service listening on :50054")
	grpcServer.Serve(lis)
}
