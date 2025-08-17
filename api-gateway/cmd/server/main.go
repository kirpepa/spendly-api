package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kirpepa/spendly-api/api-gateway/pkg/clients"
	"github.com/kirpepa/spendly-api/api-gateway/pkg/handlers"
	"github.com/kirpepa/spendly-api/api-gateway/pkg/middleware"
	"github.com/kirpepa/spendly-api/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	addrs := map[string]string{
		"auth":         cfg.AuthService,
		"user":         cfg.UserService,
		"group":        cfg.GroupService,
		"group_member": cfg.MemberService,
		"expense":      cfg.ExpenseService,
	}

	clientsGRPC, err := clients.Dial(addrs, cfg.GRPCTimeout)
	if err != nil {
		log.Fatalf("failed to dial gRPC services: %v", err)
	}
	defer clientsGRPC.Close()

	authH := &handlers.AuthHandler{Client: clientsGRPC.AuthClient, Timeout: cfg.GRPCTimeout}
	userH := &handlers.UserHandler{Client: clientsGRPC.UserClient, Timeout: cfg.GRPCTimeout}
	groupH := &handlers.GroupHandler{Client: clientsGRPC.GroupClient, Timeout: cfg.GRPCTimeout}
	gmH := &handlers.GroupMemberHandler{Client: clientsGRPC.GroupMemberClient, Timeout: cfg.GRPCTimeout}
	expH := &handlers.ExpenseHandler{Client: clientsGRPC.ExpenseClient, Timeout: cfg.GRPCTimeout}

	authMW := &middleware.AuthMiddleware{AuthClient: clientsGRPC.AuthClient, Timeout: cfg.GRPCTimeout}

	r := gin.Default()

	r.POST("/auth/register", authH.Register)
	r.POST("/auth/login", authH.Login)

	protected := r.Group("/api", authMW.Handler())

	protected.GET("/users", userH.ListUsers)
	protected.GET("/users/:id", userH.GetUser)

	protected.POST("/groups", groupH.CreateGroup)
	protected.GET("/groups/:id", groupH.GetGroup)

	protected.POST("/groups/:id/members", gmH.AddMember)
	protected.GET("/groups/:id/members", gmH.GetMembers)

	protected.POST("/expenses", expH.AddExpense)

	addr := ":" + cfg.APIPort
	log.Printf("Gateway listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run gateway: %v", err)
	}
}
