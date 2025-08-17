package handlers

import (
	"context"
	"fmt"
	"time"

	userpb "github.com/kirpepa/spendly-api/user/proto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Client  userpb.UserServiceClient
	Timeout time.Duration
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.GetUser(ctx, &userpb.GetUserRequest{UserId: parseUint64(id)})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.ListUsers(ctx, &userpb.Empty{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// helper
func parseUint64(s string) uint64 {
	var v uint64
	_, _ = fmt.Sscan(s, &v)
	return v
}
