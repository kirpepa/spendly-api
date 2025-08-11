package model

import "time"

type Expense struct {
	ID        uint `gorm:"primaryKey"`
	GroupID   uint `gorm:"index"`
	PayerID   uint
	Amount    float64
	CreatedAt time.Time
}
