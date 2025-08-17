package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	gmpb "github.com/kirpepa/spendly-api/group_member/proto"
)

type GroupMemberHandler struct {
	Client  gmpb.GroupMemberServiceClient
	Timeout time.Duration
}

type addMemberReq struct {
	GroupID uint64 `json:"group_id" binding:"required"`
	UserID  uint64 `json:"user_id" binding:"required"`
}

func (h *GroupMemberHandler) AddMember(c *gin.Context) {
	var r addMemberReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.AddMember(ctx, &gmpb.AddMemberRequest{GroupId: r.GroupID, UserId: r.UserID})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *GroupMemberHandler) GetMembers(c *gin.Context) {
	gid := parseUint64(c.Param("group_id"))
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.GetMembers(ctx, &gmpb.GetMembersRequest{GroupId: gid})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
