package service

import (
	"ads_project/model/request"
	"ads_project/model/response"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]response.ResponseAllProduct, error)
	FindById(ctx context.Context, id int) (response.ResponseProductById, error)
	Create(ctx context.Context, request request.RequestCreatedProduct) (response.ResponseCreatedProduct, error)
	Update(ctx context.Context, request request.RequestUpdatedProduct) (response.ResponseUpdatedProduct, error)
	Delete(ctx context.Context, product_id int) error
}
