package repository

import (
	"ads_project/model/domain"
	"context"
	"database/sql"
)

type ProductRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Product
	Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, id int)
}
