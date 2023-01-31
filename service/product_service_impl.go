package service

import (
	"ads_project/helper"
	"ads_project/model/domain"
	"ads_project/model/request"
	"ads_project/model/response"
	"ads_project/repository"
	"context"
	"database/sql"
	"sync"
)

type productService struct {
	db          *sql.DB
	repoProduct repository.ProductRepository
}

func NewProductService(db *sql.DB, repoProduct repository.ProductRepository) *productService {
	return &productService{
		db:          db,
		repoProduct: repoProduct,
	}
}

func (service *productService) FindAll(ctx context.Context) ([]response.ResponseAllProduct, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	product := service.repoProduct.FindAll(ctx, tx)
	responseProduct := []response.ResponseAllProduct{}
	for _, v := range product {
		responseProduct = append(responseProduct, v.ToResponseAllProduct())
	}
	return responseProduct, nil
}

func (service *productService) FindById(ctx context.Context, id int) (response.ResponseProductById, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	res := service.repoProduct.FindById(ctx, tx, id)
	defer helper.CommitOrRollback(tx)

	return res.ToResponseProductById(), nil
}

func (service *productService) Create(ctx context.Context, request request.RequestCreatedProduct) (response.ResponseCreatedProduct, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	dataProduct := domain.Product{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		dataProduct.SetProductName(&request.ProductName)
		wg.Done()
	}()

	go func() {
		dataProduct.SetDescription(&request.Description)
		wg.Done()
	}()
	wg.Wait()
	service.repoProduct.Create(ctx, tx, dataProduct)
	return dataProduct.ToResponseCreatedProduct(), err
}

func (service *productService) Update(ctx context.Context, request request.RequestUpdatedProduct) (response.ResponseUpdatedProduct, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	
	product := service.repoProduct.FindById(ctx, tx, request.ProductId)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		product.SetProductName(&request.ProductName)
		wg.Done()
	}()

	go func() {
		product.SetDescription(&request.Description)
		wg.Done()
	}()
	wg.Wait()
	
	service.repoProduct.Update(ctx, tx, product)
	return product.ToResponseUpdatedProduct(), err
}

func (service *productService) Delete(ctx context.Context, product_id int) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	service.repoProduct.Delete(ctx, tx, product_id)
	return nil
}