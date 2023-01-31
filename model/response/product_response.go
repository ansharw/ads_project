package response

import "time"

type ResponseAllProduct struct {
	ProductId   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ResponseProductById struct {
	ProductId   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ResponseCreatedProduct struct {
	// ProductId   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	// CreatedAt   time.Time `json:"created_at"`
}

type ResponseUpdatedProduct struct {
	ProductId   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	CreatedAt   string `json:"created_at"`
}
