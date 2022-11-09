package service

import (
	"context"

	"github.com/faridlan/invoice-app/app/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreate) web.OrderResponse
	Update(ctx context.Context, request web.OrderUpdate) web.OrderResponse
	Delete(ctx context.Context, orderId string)
	FindById(ctx context.Context, orderId string) web.OrderResponse
	FindAll(ctx context.Context) []web.OrderResponse
}
