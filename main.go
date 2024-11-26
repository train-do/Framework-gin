package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/train-do/Framework-gin/database"
	"github.com/train-do/Framework-gin/router"
)

func main() {
	app := gin.New()
	db, err := database.InitDB()
	if err != nil {
		return
	}
	router.APIRouter(app, db)
	fmt.Println(viper.GetString("PORT"))
	app.Run(":" + viper.GetString("PORT"))
}
func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
