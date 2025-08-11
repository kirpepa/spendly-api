package model

type GroupMember struct {
	ID      uint    `gorm:"primaryKey"`
	GroupID uint    `gorm:"index"`
	UserID  uint    `gorm:"index"`
	Balance float64 `gorm:"default:0"`
}
