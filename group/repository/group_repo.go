package repository

import (
	"github.com/kirpepa/spendly-api/group/model"
	"gorm.io/gorm"
)

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{db: db}
}

func (r *GroupRepo) Create(group *model.Group) error {
	return r.db.Create(group).Error
}

func (r *GroupRepo) GetByID(id uint) (*model.Group, error) {
	var group model.Group
	err := r.db.First(&group, id).Error
	return &group, err
}

func (r *GroupRepo) ListByOwner(ownerID uint) ([]model.Group, error) {
	var groups []model.Group
	err := r.db.Where("owner_id = ?", ownerID).Find(&groups).Error
	return groups, err
}
