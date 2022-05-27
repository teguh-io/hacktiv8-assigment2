package main

import (
	"bootcamp/hacktiv8-assigment2/controllers"
	"bootcamp/hacktiv8-assigment2/db"
	"bootcamp/hacktiv8-assigment2/repositories"
	"bootcamp/hacktiv8-assigment2/routers"
	"bootcamp/hacktiv8-assigment2/services"

	_ "bootcamp/hacktiv8-assigment2/docs"

	"github.com/gin-gonic/gin"
)

const (
	APP_PORT = ":8080"
)

// @title           Hacktiv8 Assigment 2
// @version         1.0
// @description     Simple REST API developed developed for submission assigment 2 Bootcamp at Hacktive8

// @contact.name   Teguh Ainul Darajat
// @contact.email  teguh@email.com

// @host      localhost:8080
// @BasePath  /
func main() {
	server := gin.Default()
	database := db.GetDB()

	itemRepo := repositories.NewItemRepo(database)
	itemService := services.NewItemService(itemRepo)

	orderRepo := repositories.NewOrderRepo(database)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService, itemService)
	orderRouter := routers.NewOrderRouter(server, orderController)

	swaggerRouter := routers.NewSwaggerRouter(server)

	orderRouter.Start()
	swaggerRouter.Start()
	server.Run(APP_PORT)
}
