package repositories

import (
	"bootcamp/hacktiv8-assigment2/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrders() ([]models.Order, error)
	UpdateOrder(ID int, order models.Order) error
	DeleteORder(ID int) error
}

type orderRepo struct {
	currDB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	db.AutoMigrate(&models.Order{})
	return &orderRepo{
		currDB: db,
	}
}

func (or *orderRepo) CreateOrder(order models.Order) (models.Order, error) {
	err := or.currDB.Create(&order)
	return order, err.Error
}

func (or *orderRepo) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := or.currDB.Preload("Items").Find(&orders).Error
	return orders, err
}

func (or *orderRepo) UpdateOrder(ID int, order models.Order) error {
	return or.currDB.Where("id=?", ID).Updates(&order).Error
}

func (or *orderRepo) DeleteORder(ID int) error {
	var order models.Order
	return or.currDB.Where("id=?", ID).Delete(&order).Error
}
