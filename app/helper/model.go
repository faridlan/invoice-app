package helper

import (
	"github.com/faridlan/invoice-app/app/model/domain"
	"github.com/faridlan/invoice-app/app/model/web"
)

func OrderResponse(order domain.Order) web.OrderResponse {
	orderResponse := web.OrderResponse{
		Id:        order.Id,
		OrderDate: order.OrderDate,
		CusName:   order.CusName,
		Total:     order.Total,
		Dp:        order.Dp,
		Pay:       order.Pay,
		RestOfPay: order.RestOfPay,
	}

	return orderResponse
}

func OrderResponses(orders []domain.Order) []web.OrderResponse {
	orderResponses := []web.OrderResponse{}

	for _, order := range orders {
		orderResponses = append(orderResponses, OrderResponse(order))
	}

	return orderResponses
}
