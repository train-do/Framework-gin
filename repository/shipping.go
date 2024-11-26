package repository

import (
	"github.com/train-do/Framework-gin/model"
	"gorm.io/gorm"
)

type ShippingRepository interface {
	FindAll() ([]model.Shipping, error)
}

type shippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) ShippingRepository {
	return &shippingRepository{db: db}
}

func (r *shippingRepository) FindAll() ([]model.Shipping, error) {
	var shipping []model.Shipping
	err := r.db.Find(&shipping).Error
	return shipping, err
}
