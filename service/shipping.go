package service

import (
	"github.com/train-do/Framework-gin/model"
	"github.com/train-do/Framework-gin/repository"
	"gorm.io/gorm"
)

type ShippingService interface {
	GetAllShippings() ([]model.Shipping, error)
}
type shippingService struct {
	repo repository.ShippingRepository
}

func NewShippingService(db *gorm.DB) ShippingService {
	return &shippingService{
		repo: repository.NewShippingRepository(db),
	}
}

func (s *shippingService) GetAllShippings() ([]model.Shipping, error) {
	return s.repo.FindAll()
}
