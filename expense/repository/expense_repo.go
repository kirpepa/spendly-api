package repository

import (
	"github.com/kirpepa/spendly-api/expense/model"
	"gorm.io/gorm"
)

type ExpenseRepo struct {
	db *gorm.DB
}

func NewExpenseRepo(db *gorm.DB) *ExpenseRepo {
	return &ExpenseRepo{db: db}
}

func (r *ExpenseRepo) Create(expense *model.Expense) error {
	return r.db.Create(expense).Error
}
