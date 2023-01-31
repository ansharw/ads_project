package request

type RequestCreatedProduct struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}

type RequestUpdatedProduct struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}

type RequestDeletedProduct struct {
	ProductId int `json:"product_id"`
}
