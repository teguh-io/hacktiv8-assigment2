package main

import (
	"bootcamp/hacktiv8-assigment2/controllers"
	"bootcamp/hacktiv8-assigment2/db"
	"bootcamp/hacktiv8-assigment2/repositories"
	"bootcamp/hacktiv8-assigment2/routers"
	"bootcamp/hacktiv8-assigment2/services"

	"github.com/gin-gonic/gin"
)

const (
	APP_PORT = ":8080"
)

func main() {
	server := gin.Default()
	database := db.GetDB()

	itemRepo := repositories.NewItemRepo(database)
	itemService := services.NewItemService(itemRepo)

	orderRepo := repositories.NewOrderRepo(database)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService, itemService)
	orderRouter := routers.NewOrderRouter(server, orderController)

	orderRouter.Start()
	server.Run(APP_PORT)
}
