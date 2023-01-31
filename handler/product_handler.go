package handler

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	FindAllProduct(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
