package controller

import (
	"net/http"

	"github.com/faridlan/invoice-app/app/helper"
	"github.com/faridlan/invoice-app/app/model/web"
	"github.com/faridlan/invoice-app/app/service"
	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderCreate := web.OrderCreate{}
	helper.ReadFromRequestBody(request, &orderCreate)

	order := controller.OrderService.Create(request.Context(), orderCreate)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   order,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderUpdate := web.OrderUpdate{}
	helper.ReadFromRequestBody(request, &orderUpdate)

	orderUpdate.Id = params.ByName("Id")

	order := controller.OrderService.Update(request.Context(), orderUpdate)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   order,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("Id")

	controller.OrderService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("Id")

	order := controller.OrderService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   order,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	orders := controller.OrderService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orders,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
