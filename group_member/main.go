package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/kirpepa/spendly-api/group_member/proto"
	"github.com/kirpepa/spendly-api/group_member/repository"
	"github.com/kirpepa/spendly-api/group_member/service"

	"github.com/kirpepa/spendly-api/internal/config"
	"github.com/kirpepa/spendly-api/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	database := db.InitDB(cfg.DBUrl)

	repo := repository.NewGroupMemberRepo(database)
	server := service.NewGroupMemberServer(repo)

	grpcServer := grpc.NewServer()
	proto.RegisterGroupMemberServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("GroupMember service listening on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
