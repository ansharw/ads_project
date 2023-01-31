package repository

import (
	"ads_project/model/domain"
	"context"
	"database/sql"
	"fmt"
)

type productRepository struct {
}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (repo *productRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	query := "SELECT product_id, product_name, product_description, created_at FROM product ORDER BY product_id DESC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		rows.Scan(product.GetProductId(), product.GetProductName(), product.GetDescription(), product.GetCreatedAt())
		products = append(products, product)
	}
	return products
}

func (repo *productRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Product {
	query := "SELECT product_id, product_name, product_description, created_at FROM product WHERE product_id = ?"
	product := domain.Product{}
	row := tx.QueryRowContext(ctx, query, id)
	row.Scan(product.GetProductId(), product.GetProductName(), product.GetDescription(), product.GetCreatedAt())
	return product
}

func (repo *productRepository) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := "INSERT INTO product(product_name, product_description, created_at) VALUES (?, ?, ?)"
	res, err := tx.ExecContext(ctx, query, product.GetProductName(), product.GetDescription(), product.GetCreatedAt())
	if err != nil {
		panic(err)
	}
	lastId, _ := res.LastInsertId()
	id := int(lastId)
	fmt.Println("ini id dari insert ", id)
	product.SetProductId(&id)
	fmt.Println("product id: ", *product.GetProductId())
	return product
}

func (repo *productRepository) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := `UPDATE product SET product_name = ?, product_description = ? WHERE product_id = ?`
	_, err := tx.ExecContext(ctx, query, product.GetProductName(), product.GetDescription(), product.GetProductId())
	if err != nil {
		panic(err)
	}
	return product
}

func (repo *productRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM product WHERE product_id = ?"
	_, err := tx.ExecContext(ctx, query, id)

	if err != nil {
		panic(err)
	}
}
