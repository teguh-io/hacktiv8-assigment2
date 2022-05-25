package repositories

import (
	"bootcamp/hacktiv8-assigment2/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(item models.Item) error
	GetItems() ([]models.Item, error)
	UpdateItem(ID int, item models.Item) error
	DeleteItem(ID int) error
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

func (ir *itemRepo) CreateItem(item models.Item) error {
	return ir.currDB.Create(&item).Error
}

func (ir *itemRepo) GetItems() ([]models.Item, error) {
	var items []models.Item
	err := ir.currDB.Find(&items).Error
	return items, err
}

func (ir *itemRepo) UpdateItem(ID int, item models.Item) error {
	return ir.currDB.Where("id=?", ID).Updates(&item).Error
}

func (ir *itemRepo) DeleteItem(ID int) error {
	var item models.Item
	return ir.currDB.Where("id=?", ID).Delete(&item).Error
}
