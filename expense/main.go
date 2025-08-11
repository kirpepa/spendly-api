package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/kirpepa/spendly-api/expense/proto"
	"github.com/kirpepa/spendly-api/expense/repository"
	"github.com/kirpepa/spendly-api/expense/service"

	gmRepo "github.com/kirpepa/spendly-api/group_member/repository"

	"github.com/kirpepa/spendly-api/internal/config"
	"github.com/kirpepa/spendly-api/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	database := db.InitDB(cfg.DBUrl)

	expRepo := repository.NewExpenseRepo(database)
	gmRepository := gmRepo.NewGroupMemberRepo(database)

	server := service.NewExpenseServer(expRepo, gmRepository)

	grpcServer := grpc.NewServer()
	proto.RegisterExpenseServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Expense service listening on :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
