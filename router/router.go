package router

import (
	"github.com/gin-gonic/gin"
	"github.com/train-do/Framework-gin/controller"
	"gorm.io/gorm"
)

func APIRouter(router *gin.Engine, db *gorm.DB) {
	ctl := controller.NewController(db)
	router.GET("/api/shippings", ctl.Shipping.GetShippings)
	router.POST("/api/shippings", ctl.Shipping.GetOngkir)
}
