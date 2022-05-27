package services

import (
	"bootcamp/hacktiv8-assigment2/models"
	"bootcamp/hacktiv8-assigment2/params"
	"bootcamp/hacktiv8-assigment2/repositories"
)

type ItemService interface {
	CreateItem(orderID int, request []params.Item) error
	UpdateItemsByOrderID(ID int, request []params.Item) error
	DeleteItem(ID int) error
}

type itemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(ir repositories.ItemRepo) ItemService {
	return &itemService{
		itemRepo: ir,
	}
}

func toParamItems(itemModels []models.Item) []params.Item {
	itemParams := make([]params.Item, len(itemModels))
	for idx, itemModel := range itemModels {
		itemParams[idx] = toParamItem(itemModel)
	}

	return itemParams
}

func toParamItem(itemModel models.Item) params.Item {
	return params.Item{
		ItemCode:    &itemModel.ItemCode,
		Description: &itemModel.Description,
		Quantity:    &itemModel.Quantitiy,
	}
}

func toItemModels(orderID int, itemParams []params.Item) []models.Item {
	itemModels := make([]models.Item, len(itemParams))
	for idx, itemParam := range itemParams {
		itemModels[idx] = toItemModel(orderID, itemParam)
	}

	return itemModels
}

func toItemModel(orderID int, itemParam params.Item) models.Item {
	return models.Item{
		ItemCode:    *itemParam.ItemCode,
		Description: *itemParam.Description,
		Quantitiy:   *itemParam.Quantity,
		OrderID:     orderID,
	}
}

func (is *itemService) CreateItem(orderID int, request []params.Item) error {
	itemModels := toItemModels(orderID, request)

	return is.itemRepo.CreateItem(itemModels)
}

func (is *itemService) UpdateItemsByOrderID(ID int, request []params.Item) error {
	itemModels := toItemModels(ID, request)

	return is.itemRepo.UpdateItemByOrderID(ID, itemModels)
}

func (is *itemService) DeleteItem(ID int) error {
	panic("asd")
}
