package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/faridlan/invoice-app/app/helper"
	"github.com/faridlan/invoice-app/app/model/domain"
	"github.com/faridlan/invoice-app/app/model/web"
	"github.com/faridlan/invoice-app/app/repository"
	"github.com/google/uuid"
)

type OrderServiceImpl struct {
	OrderRepo repository.OrderRepository
	DB        *sql.DB
}

func (Service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreate) web.OrderResponse {
	tx, err := Service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		Id:        uuid.NewString(),
		OrderDate: time.Now().Unix(),
		CusName:   request.CusName,
		Total:     request.Total,
		Dp:        request.Dp,
		Pay:       request.Pay,
		RestOfPay: request.RestOfPay,
	}

	result, err := Service.OrderRepo.Create(ctx, tx, order)
	helper.PanicIfErr(err)

	return helper.OrderResponse(*result)
}

func (Service *OrderServiceImpl) Update(ctx context.Context, request web.OrderUpdate) web.OrderResponse {
	tx, err := Service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	order, err := Service.OrderRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	order.CusName = request.CusName
	order.Total = request.Total
	order.Dp = request.Dp
	order.Pay = request.Pay
	order.RestOfPay = request.RestOfPay

	result, err := Service.OrderRepo.Update(ctx, tx, *order)
	helper.PanicIfErr(err)

	return helper.OrderResponse(*result)
}

func (Service *OrderServiceImpl) Delete(ctx context.Context, orderId string) {
	tx, err := Service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	order, err := Service.OrderRepo.FindById(ctx, tx, orderId)
	helper.PanicIfErr(err)

	err = Service.OrderRepo.Delete(ctx, tx, order)
	helper.PanicIfErr(err)
}

func (Service *OrderServiceImpl) FindById(ctx context.Context, orderId string) web.OrderResponse {
	tx, err := Service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	order, err := Service.OrderRepo.FindById(ctx, tx, orderId)
	helper.PanicIfErr(err)

	return helper.OrderResponse(*order)
}

func (Service *OrderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {
	tx, err := Service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	order, err := Service.OrderRepo.FindAll(ctx, tx)
	helper.PanicIfErr(err)

	return helper.OrderResponses(order)
}
