package controllers

import (
	"bootcamp/hacktiv8-assigment2/helpers"
	"bootcamp/hacktiv8-assigment2/params"
	"bootcamp/hacktiv8-assigment2/services"
	"fmt"
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
		msg := fmt.Sprintf("Failed to create order: %s", err.Error())
		jsonResponse = helpers.JsonResponse(http.StatusNotAcceptable, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.itemService.CreateItem(order.ID, request.Items)
	if err != nil {
		msg := fmt.Sprintf("Failed to create item order: %s", err.Error())
		jsonResponse = helpers.JsonResponse(http.StatusNotAcceptable, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helpers.JsonResponse(http.StatusOK, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

func (or *orderController) GetOrders(ctx *gin.Context) {
	var jsonResponse params.Response
	orders, err := or.orderService.GetOrders()
	if err != nil {
		msg := fmt.Sprintf("Failed to get orders data: %s", err.Error())
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helpers.JsonResponse(http.StatusOK, nil, orders)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

func (or *orderController) UpdateOrder(ctx *gin.Context) {
	var request params.Order
	var jsonResponse params.Response

	IDstr := ctx.Param("id")
	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		msg := "something is wrong with the params given"
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		msg := "something is wrong with the body request given"
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.UpdateOrder(ID, request)
	if err != nil {
		msg := fmt.Sprintf("Failed to update order: %s", err.Error())
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	if request.Items != nil {
		err = or.itemService.UpdateItemsByOrderID(ID, request.Items)
		if err != nil {
			msg := fmt.Sprintf("Failed to update order items: %s", err.Error())
			jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
			ctx.JSON(jsonResponse.Status, jsonResponse)
			return
		}
	}

	jsonResponse = helpers.JsonResponse(http.StatusOK, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

func (or *orderController) DeleteOrder(ctx *gin.Context) {
	var jsonResponse params.Response
	IDStr := ctx.Param("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		msg := "something is wrong with the params given"
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.DeleteOrder(ID)
	if err != nil {
		msg := fmt.Sprintf("Failed to delete order: %s", err.Error())
		jsonResponse = helpers.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helpers.JsonResponse(http.StatusOK, nil, "the data has been deleted")
	ctx.JSON(jsonResponse.Status, jsonResponse)
}
