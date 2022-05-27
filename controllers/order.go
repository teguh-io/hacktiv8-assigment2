package controllers

import (
	"bootcamp/hacktiv8-assigment2/helpers"
	"bootcamp/hacktiv8-assigment2/params"
	"bootcamp/hacktiv8-assigment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderController struct {
	orderService services.OrderService
	itemService  services.ItemService
}

func NewOrderController(os services.OrderService, is services.ItemService) OrderController {
	return &orderController{
		orderService: os,
		itemService:  is,
	}
}

// @Create Order
// @tag.name Order
// @Description API for create order
// @Accept json
// @Produce json
// @Tags Orders
// @param Body body params.Order true "Create order"
// @Success 200 {object} params.OrderResponse
// @Router /orders [POST]
func (or *orderController) CreateOrder(ctx *gin.Context) {
	var request params.Order
	var jsonResponse params.Response
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order, err := or.orderService.CreateOrder(*request.CustomerName)
	if err != nil {
		msg := "Failed to create order"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusInternalServerError, &msg, &additionalInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.itemService.CreateItem(order.ID, request.Items)
	if err != nil {
		msg := "Failed to create item order"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusInternalServerError, &msg, &additionalInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	msg := "order successfully created"
	jsonResponse = helpers.JsonResponse(http.StatusCreated, &msg, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

// @Get Order
// @tag.name Order
// @Description API for Get All order data
// @Produce json
// @Tags Orders
// @Success 200 {object} params.OrderResponse
// @Router /orders [GET]
func (or *orderController) GetOrders(ctx *gin.Context) {
	var jsonResponse params.Response
	orders, err := or.orderService.GetOrders()
	if err != nil {
		msg := "Failed to get orders"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusInternalServerError, &msg, &additionalInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helpers.JsonResponse(http.StatusOK, nil, nil, orders)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

// @Update Order
// @tag.name Order
// @Description API for update order data
// @Accept json
// @Produce json
// @Tags Orders
// @param Body body params.Order true "Update Order"
// @Success 200 {object} params.OrderResponse
// @Router /orders/:id [PUT]
func (or *orderController) UpdateOrder(ctx *gin.Context) {
	var request params.Order
	var jsonResponse params.Response

	IDstr := ctx.Param("id")
	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		msg := "something is wrong with the params given"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, &additionalInfo, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		msg := "something is wrong with the body request given"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, &additionalInfo, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.UpdateOrder(ID, request)
	if err != nil {
		msg := "Failed to update order"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, &additionalInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	if request.Items != nil {
		err = or.itemService.UpdateItemsByOrderID(ID, request.Items)
		if err != nil {
			msg := "Failed to update order items"
			additionalInfo := err.Error()
			jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, &additionalInfo, nil)
			ctx.JSON(jsonResponse.Status, jsonResponse)
			return
		}
	}

	msg := "order successfully updated"
	jsonResponse = helpers.JsonResponse(http.StatusOK, &msg, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

// @Delete Order
// @tag.name Order
// @Description API for Delete order data
// @Produce json
// @Tags Orders
// @Success 200 {object} params.OrderResponse
// @Router /orders/:id [DELETE]
func (or *orderController) DeleteOrder(ctx *gin.Context) {
	var jsonResponse params.Response
	IDStr := ctx.Param("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		msg := "something is wrong with the params given"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, &additionalInfo, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.DeleteOrder(ID)
	if err != nil {
		msg := "Failed to delete order"
		additionalInfo := err.Error()
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, &additionalInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	msg := "the order successfully deleted"
	jsonResponse = helpers.JsonResponse(http.StatusOK, &msg, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}
