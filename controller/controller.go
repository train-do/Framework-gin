package controller

import "gorm.io/gorm"

type Controller struct {
	Shipping ShippingController
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Shipping: *NewShippingController(db),
	}
}
