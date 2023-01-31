package domain

import (
	"ads_project/model/response"
	"time"
)

type Product struct {
	productId   int
	productName string
	description string
	createdAt   time.Time
}

func (c *Product) ToResponseAllProduct() response.ResponseAllProduct {
	return response.ResponseAllProduct{
		ProductId:   c.productId,
		ProductName: c.productName,
		Description: c.description,
		CreatedAt:   c.createdAt.Format("2006-01-02 15:04:05"),
	}
}

func (c *Product) ToResponseProductById() response.ResponseProductById {
	return response.ResponseProductById{
		ProductId:   c.productId,
		ProductName: c.productName,
		Description: c.description,
		CreatedAt:   c.createdAt,
	}
}

func (c *Product) ToResponseCreatedProduct() response.ResponseCreatedProduct {
	return response.ResponseCreatedProduct{
		ProductId:   c.productId,
		ProductName: c.productName,
		Description: c.description,
		CreatedAt:   c.createdAt.Format("2006-01-02 15:04:05"),
	}
}

func (c *Product) ToResponseUpdatedProduct() response.ResponseUpdatedProduct {
	return response.ResponseUpdatedProduct{
		ProductId:   c.productId,
		ProductName: c.productName,
		Description: c.description,
		CreatedAt:   c.createdAt.Format("2006-01-02 15:04:05"),
	}
}

func (c *Product) SetProduct(productId *int, productName, description *string, createdAt *time.Time) {
	c.productId = *productId
	c.productName = *productName
	c.description = *description
	c.createdAt = *createdAt
}

func (c *Product) SetProductId(productId *int) {
	c.productId = *productId
}

func (c *Product) SetProductName(productName *string) {
	c.productName = *productName
}

func (c *Product) SetDescription(description *string) {
	c.description = *description
}

func (c *Product) SetCreatedAt(createdAt *time.Time) {
	c.createdAt = *createdAt
}

func (c *Product) GetProduct() (*int, *string, *string, *time.Time) {
	return &c.productId, &c.productName, &c.description, &c.createdAt
}

func (c *Product) GetProductId() *int {
	return &c.productId
}

func (c *Product) GetProductName() *string {
	return &c.productName
}

func (c *Product) GetDescription() *string {
	return &c.description
}

func (c *Product) GetCreatedAt() *time.Time {
	return &c.createdAt
}

func NewProductDomain(productId int, productName, description string, createdAt time.Time) *Product {
	return &Product{
		productId:   productId,
		productName: productName,
		description: description,
		createdAt:   createdAt,
	}
}
