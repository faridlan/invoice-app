package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/invoice-app/app/model/domain"
)

type OrderRepositoryImpl struct {
}

func (repository *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order domain.Order) (*domain.Order, error) {
	SQL := "insert into orders(id, order_date, cus_name, total, dp, pay, rest_of_pay) values (?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, order.Id, order.OrderDate, order.CusName, order.Total, order.Dp, order.Pay, order.RestOfPay)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order domain.Order) (*domain.Order, error) {
	SQL := "update orders set cus_name = ?, total = ?, dp = ?, pay = ?, rest_of_pay = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.CusName, order.Total, order.Dp, order.Pay, order.RestOfPay, order.Id)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, order domain.Order) error {
	SQL := "delet from orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId string) (*domain.Order, error) {

	SQL := "Select id, order_date, cus_name, total, dp, pay, rest_of_pay from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order domain.Order
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.OrderDate, &order.CusName, &order.Total, &order.Dp, &order.Pay, &order.RestOfPay)
		if err != nil {
			return nil, err
		}

		return &order, nil
	} else {
		return nil, errors.New("order not found")
	}
}
func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Order, error) {
	SQL := "Select id, order_date, cus_name, total, dp, pay, rest_of_pay from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		var order domain.Order

		err := rows.Scan(&order.Id, &order.OrderDate, &order.CusName, &order.Total, &order.Dp, &order.Pay, &order.RestOfPay)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
