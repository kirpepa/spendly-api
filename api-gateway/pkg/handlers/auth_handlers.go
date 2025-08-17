package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	authpb "github.com/kirpepa/spendly-api/auth/proto"
)

type AuthHandler struct {
	Client  authpb.AuthServiceClient
	Timeout time.Duration
}

type registerReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var r registerReq

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	regReq := &authpb.RegisterRequest{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}

	resp, err := h.Client.Register(ctx, regReq)

	if err != nil || resp.Error != "" {
		c.JSON(500, gin.H{"error": resp.Error})
		return
	}
	c.JSON(201, gin.H{"token": resp.Token})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var r loginReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.Login(ctx, &authpb.LoginRequest{
		Email:    r.Email,
		Password: r.Password,
	})
	if err != nil || resp.Error != "" {
		c.JSON(401, gin.H{"error": resp.Error})
		return
	}
	c.JSON(200, gin.H{"token": resp.Token})
}
