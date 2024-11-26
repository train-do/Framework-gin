package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/train-do/Framework-gin/model"
	"github.com/train-do/Framework-gin/service"
	"gorm.io/gorm"
)

type ShippingController struct {
	service service.ShippingService
}

func NewShippingController(db *gorm.DB) *ShippingController {
	return &ShippingController{
		service: service.NewShippingService(db),
	}
}

func (c *ShippingController) GetShippings(ctx *gin.Context) {
	shippings, err := c.service.GetAllShippings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shippings)
}
func (c *ShippingController) GetOngkir(ctx *gin.Context) {
	client := http.Client{}
	var ongkir model.Ongkir
	if err := ctx.ShouldBindJSON(&ongkir); err != nil {
		ctx.String(http.StatusBadRequest, "Error: %s", err.Error())
		return
	}
	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s;%s?overview=false", ongkir.LatLongOrigin, ongkir.LatLongDestination)
	// url := `https://router.project-osrm.org/route/v1/driving/107.8276122,-6.199131;106.949822,-6.2345118?overview=false`
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// fmt.Printf("%+v\n", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	// fmt.Printf("%+v<<<<\n", resp)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	routes, _ := result["routes"].([]interface{})
	route, _ := routes[0].(map[string]interface{})
	distance, _ := route["distance"].(float64)
	ongkir.Distance = distance
	ongkir.Price = int(distance / 1000.0 * 2000)
	ctx.JSON(http.StatusOK, routes)
}
