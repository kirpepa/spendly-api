package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	expbp "github.com/kirpepa/spendly-api/expense/proto"
)

type ExpenseHandler struct {
	Client  expbp.ExpenseServiceClient
	Timeout time.Duration
}

type addExpenseReq struct {
	GroupID uint64  `json:"group_id" binding:"required"`
	PayerID uint64  `json:"payer_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

func (h *ExpenseHandler) AddExpense(c *gin.Context) {
	var r addExpenseReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	resp, err := h.Client.AddExpense(ctx, &expbp.AddExpenseRequest{
		GroupId: r.GroupID,
		PayerId: r.PayerID,
		Amount:  r.Amount,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
