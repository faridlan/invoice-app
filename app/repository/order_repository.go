package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/invoice-app/app/model/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *sql.Tx, order domain.Order) (*domain.Order, error)
	Update(ctx context.Context, tx *sql.Tx, order domain.Order) (*domain.Order, error)
	Delete(ctx context.Context, tx *sql.Tx, order domain.Order) error
	FindById(ctx context.Context, tx *sql.Tx, orderId string) (*domain.Order, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Order, error)
}
