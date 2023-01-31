package main

import (
	"ads_project/database"
	"ads_project/handler"
	"ads_project/repository"
	"ads_project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	db := database.GetConnection()

	repoProduct := repository.NewProductRepository()
	serviceProduct := service.NewProductService(db, repoProduct)
	handlerProduct := handler.NewProductHandler(serviceProduct)

	api.GET("/health")
	api.GET("/time")

	api.GET("/product", handlerProduct.FindAllProduct)
	api.POST("/product", handlerProduct.Create)
	api.PUT("/product", handlerProduct.Update)
	api.DELETE("/product/:product_id", handlerProduct.Delete)

	r.Run("localhost:9000")
}
