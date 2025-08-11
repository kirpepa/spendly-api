package service

import (
	"context"
	"fmt"
	"github.com/kirpepa/spendly-api/expense/model"
	"github.com/kirpepa/spendly-api/expense/proto"
	"github.com/kirpepa/spendly-api/expense/repository"

	gmRepo "github.com/kirpepa/spendly-api/group_member/repository"
)

type ExpenseServer struct {
	proto.UnimplementedExpenseServiceServer
	expenseRepo     *repository.ExpenseRepo
	groupMemberRepo *gmRepo.GroupMemberRepo
}

func NewExpenseServer(expRepo *repository.ExpenseRepo, gmRepo *gmRepo.GroupMemberRepo) *ExpenseServer {
	return &ExpenseServer{expenseRepo: expRepo, groupMemberRepo: gmRepo}
}

func (s *ExpenseServer) AddExpense(ctx context.Context, req *proto.AddExpenseRequest) (*proto.ExpenseResponse, error) {
	exp := &model.Expense{
		GroupID: uint(req.GroupId),
		PayerID: uint(req.PayerId),
		Amount:  req.Amount,
	}
	if err := s.expenseRepo.Create(exp); err != nil {
		return nil, err
	}

	members, err := s.groupMemberRepo.GetByGroup(uint(req.GroupId))
	if err != nil {
		return nil, err
	}
	if len(members) == 0 {
		return nil, fmt.Errorf("no members in group")
	}

	share := req.Amount / float64(len(members))

	for _, m := range members {
		if m.UserID == uint(req.PayerId) {
			s.groupMemberRepo.UpdateBalance(m.GroupID, m.UserID, req.Amount-share)
		} else {
			s.groupMemberRepo.UpdateBalance(m.GroupID, m.UserID, -share)
		}
	}

	return &proto.ExpenseResponse{Message: "Expense added and balances updated"}, nil
}
