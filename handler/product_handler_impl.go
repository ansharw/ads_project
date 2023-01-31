package handler

import (
	"ads_project/model/api"
	"ads_project/model/request"
	"ads_project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (handler *productHandler) FindAllProduct(c *gin.Context) {
	AllProduct, err := handler.productService.FindAll(c)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   AllProduct,
	})
}

func (handler *productHandler) Create(c *gin.Context) {
	var req request.RequestCreatedProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ProductCreated, err := handler.productService.Create(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(ProductCreated.ProductId)
	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "Successfully create product data",
		Data:   ProductCreated,
	})
}

func (handler *productHandler) Update(c *gin.Context) {
	var req request.RequestUpdatedProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ProductUpdated, err := handler.productService.Update(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "Successfully update product data",
		Data:   ProductUpdated,
	})
}

// Delete by id di requestbody
// func (handler *productHandler) Delete(c *gin.Context) {
// 	var req request.RequestDeletedProduct

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := handler.productService.Delete(c, req.ProductId)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, api.ApiResponse{
// 		Code:   200,
// 		Status: "Successfully delete product data",
// 	})
// }

// Delete by product id on parameter
func (handler *productHandler) Delete(c *gin.Context) {
	val, ok := c.Params.Get("product_id")
	if !ok {
		panic("product id not found")
	}

	product_id, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	err = handler.productService.Delete(c, product_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "Successfully delete product data",
	})
}
