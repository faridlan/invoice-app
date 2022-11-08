package mock

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/invoice-app/app/helper"
	"github.com/faridlan/invoice-app/app/model/domain"
	"github.com/faridlan/invoice-app/app/repository"
	"github.com/stretchr/testify/assert"
)

var i = &domain.Order{
	Id:        "INVNOV001",
	OrderDate: 1667920980102,
	CusName:   "Udin",
	Total:     3000000,
	Dp:        1000000,
	Pay:       500000,
	RestOfPay: 1500000,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock

}

func commitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		helper.PanicIfErr(errRollback)
		panic(err)
	} else {
		err := tx.Commit()
		helper.PanicIfErr(err)
	}
}

func TestFindById(t *testing.T) {
	db, mock := NewMock()
	tx, err := db.Begin()
	helper.PanicIfErr(err)
	defer commitOrRollback(tx)
	defer db.Close()

	repo := repository.OrderRepositoryImpl{}

	mock.ExpectBegin()
	query := "SELECT id, order_date, cus_name, total, dp, pay, rest_of_pay FROM orders WHERE id = \\?"
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"id", "order_date", "cus_name", "total", "dp", "pay", "rest_of_pay"})

	mock.ExpectQuery(query).WithArgs(i.Id).WillReturnRows(rows)

	order, err := repo.FindById(context.Background(), tx, i.Id)
	assert.NotNil(t, order)
	assert.NoError(t, err)
}
