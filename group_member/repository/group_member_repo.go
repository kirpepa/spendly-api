package repository

import (
	"github.com/kirpepa/spendly-api/group_member/model"
	"gorm.io/gorm"
)

type GroupMemberRepo struct {
	db *gorm.DB
}

func NewGroupMemberRepo(db *gorm.DB) *GroupMemberRepo {
	return &GroupMemberRepo{db: db}
}

func (r *GroupMemberRepo) AddMember(member *model.GroupMember) error {
	return r.db.Create(member).Error
}

func (r *GroupMemberRepo) GetByGroup(groupID uint) ([]model.GroupMember, error) {
	var members []model.GroupMember
	err := r.db.Where("group_id = ?", groupID).Find(&members).Error
	return members, err
}

func (r *GroupMemberRepo) UpdateBalance(groupID, userID uint, delta float64) error {
	return r.db.Model(&model.GroupMember{}).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Update("balance", gorm.Expr("balance + ?", delta)).Error
}
