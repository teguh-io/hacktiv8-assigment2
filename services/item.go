package services

import (
	"bootcamp/hacktiv8-assigment2/models"
	"bootcamp/hacktiv8-assigment2/params"
	"bootcamp/hacktiv8-assigment2/repositories"
)

type ItemService interface {
	CreateItem(request params.Item, orderID int) error
	UpdateItem(ID int, request params.Item) error
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

func (is *itemService) CreateItem(request params.Item, orderID int) error {
	itemModel := models.Item{
		ItemCode:    *request.ItemCode,
		Description: *request.Description,
		Quantitiy:   *request.Quantity,
		OrderID:     orderID,
	}

	return is.itemRepo.CreateItem(itemModel)
}

func (is *itemService) UpdateItem(ID int, request params.Item) error {
	var itemModel models.Item
	if request.ItemCode != nil {
		itemModel.ItemCode = *request.ItemCode
	}
	if request.Description != nil {
		itemModel.Description = *request.Description
	}
	if request.Quantity != nil {
		itemModel.Quantitiy = *request.Quantity
	}

	return is.itemRepo.UpdateItem(ID, itemModel)
}

func (is *itemService) DeleteItem(ID int) error {
	return is.itemRepo.DeleteItem(ID)
}
