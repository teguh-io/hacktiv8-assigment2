package repositories

import (
	"bootcamp/hacktiv8-assigment2/models"
	"bootcamp/hacktiv8-assigment2/params"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(item []models.Item) error
	GetItems() ([]models.Item, error)
	UpdateItemByOrderID(orderId int, item []models.Item) error
}

type itemRepo struct {
	currDB *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	db.AutoMigrate(&models.Item{})
	return &itemRepo{
		currDB: db,
	}
}

func (ir *itemRepo) CreateItem(item []models.Item) error {
	return ir.currDB.Create(&item).Error
}

func (ir *itemRepo) GetItems() ([]models.Item, error) {
	var items []models.Item
	err := ir.currDB.Find(&items).Error
	return items, err
}

func (ir *itemRepo) UpdateItemByOrderID(orderID int, items []models.Item) error {
	var item params.Item
	err := ir.currDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("order_id=?", orderID).Delete(&item).Error
		if err != nil {
			return err
		}

		err = tx.Create(&items).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (ir *itemRepo) DeleteItem(ID int) error {
	var item models.Item
	return ir.currDB.Where("id=?", ID).Delete(&item).Error
}
