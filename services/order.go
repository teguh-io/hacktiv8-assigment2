package services

import (
	"bootcamp/hacktiv8-assigment2/models"
	"bootcamp/hacktiv8-assigment2/params"
	"bootcamp/hacktiv8-assigment2/repositories"
	"time"
)

type OrderService interface {
	CreateOrder(customerName string) error
	GetOrders() ([]params.Order, error)
	UpdateOrder(ID int, request params.Order) error
	DeleteOrder(ID int) error
}

type orderService struct {
	orderRepo repositories.OrderRepo
}

func NewOrderService(or repositories.OrderRepo) OrderService {
	return &orderService{
		orderRepo: or,
	}
}

func toOrderParams(orderModels []models.Order) []params.Order {
	orderParams := make([]params.Order, len(orderModels))
	for idx, orderModel := range orderModels {
		orderParams[idx] = toOrderParam(orderModel)
	}

	return orderParams
}

func toOrderParam(orderModel models.Order) params.Order {
	return params.Order{
		CustomerName: &orderModel.CustomerName,
		Items:        toParamItems(orderModel.Items),
		OrderedAt:    &orderModel.OrderedAt,
	}
}

func (os *orderService) CreateOrder(customerName string) error {
	order := models.Order{
		OrderedAt:    time.Now(),
		CustomerName: customerName,
	}

	return os.orderRepo.CreateOrder(order)
}

func (os *orderService) GetOrders() ([]params.Order, error) {
	orders, err := os.orderRepo.GetOrders()
	if err != nil {
		return nil, err
	}

	orderParams := toOrderParams(orders)
	return orderParams, nil
}

func (os *orderService) UpdateOrder(ID int, request params.Order) error {
	var order models.Order
	if request.CustomerName != nil {
		order.CustomerName = *request.CustomerName
	}
	if request.OrderedAt != nil {
		order.OrderedAt = *request.OrderedAt
	}

	return os.orderRepo.UpdateOrder(ID, order)
}

func (os *orderService) DeleteOrder(ID int) error {
	return os.orderRepo.DeleteORder(ID)
}
