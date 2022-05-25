package db

import (
	"bootcamp/hacktiv8-assigment2/configs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	config := configs.GetConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
