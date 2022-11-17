package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/invoice-app/app/model/domain"
	"github.com/faridlan/invoice-app/app/repository"
	"github.com/stretchr/testify/assert"
)

var order = domain.Order{
	Id:        "INVNOV001",
	OrderDate: 1667920980102,
	CusName:   "Udin",
	Total:     3000000,
	Dp:        1000000,
	Pay:       500000,
	RestOfPay: 1500000,
}

var orderRepository = repository.NewOrderRepository()

func TestFindByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	orders := []string{"id", "order_date", "cus_name", "total", "dp", "pay", "rest_of_pay"}
	rows := sqlmock.NewRows(orders).
		AddRow(order.Id, order.OrderDate, order.CusName, order.Total, order.Dp, order.Pay, order.RestOfPay)

	eq := mock.ExpectQuery("Select id, order_date, cus_name, total, dp, pay, rest_of_pay from orders where id = ?").WithArgs(order.Id).WillReturnRows(rows)

	result, err := orderRepository.FindById(context.Background(), db, order.Id)
	if err != nil {
		t.Errorf("result error %s", err)
	}

	assert.Nil(t, err)
	fmt.Println(eq)
	fmt.Println(result)
}
