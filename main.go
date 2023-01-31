package main

import (
	"ads_project/database"
	"ads_project/handler"
	"ads_project/repository"
	"ads_project/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	db := database.GetConnection()

	repoProduct := repository.NewProductRepository()
	serviceProduct := service.NewProductService(db, repoProduct)
	handlerProduct := handler.NewProductHandler(serviceProduct)

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	api.GET("/time", func(c *gin.Context) {
		var now = time.Now()
		c.JSON(http.StatusOK, gin.H{
			"time":    fmt.Sprintf("Tanggal: %v %v %v Waktu: %v.%v.%v", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second()),
			"message": "Hello, World!",
		})
	})

	api.GET("/product", handlerProduct.FindAllProduct)
	api.POST("/product", handlerProduct.Create)
	api.PUT("/product", handlerProduct.Update)
	api.DELETE("/product/:product_id", handlerProduct.Delete)

	r.Run("localhost:9000")
}
