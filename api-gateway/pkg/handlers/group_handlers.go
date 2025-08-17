package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	grouppb "github.com/kirpepa/spendly-api/group/proto"
)

type GroupHandler struct {
	Client  grouppb.GroupServiceClient
	Timeout time.Duration
}

type createGroupReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerID     uint64 `json:"owner_id" binding:"required"`
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var r createGroupReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.CreateGroup(ctx, &grouppb.CreateGroupRequest{
		Name:        r.Name,
		Description: r.Description,
		OwnerId:     r.OwnerID,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, resp)
}

func (h *GroupHandler) GetGroup(c *gin.Context) {
	id := parseUint64(c.Param("id"))
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.GetGroup(ctx, &grouppb.GetGroupRequest{GroupId: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
