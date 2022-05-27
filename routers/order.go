package routers

import (
	"bootcamp/hacktiv8-assigment2/controllers"

	"github.com/gin-gonic/gin"
)

type orderRouter struct {
	router          *gin.Engine
	orderController controllers.OrderController
}

func NewOrderRouter(r *gin.Engine, oc controllers.OrderController) *orderRouter {
	return &orderRouter{
		router:          r,
		orderController: oc,
	}

}

func (or *orderRouter) Start() {
	or.router.POST("/orders", or.orderController.CreateOrder)
	or.router.GET("/orders", or.orderController.GetOrders)
	or.router.PUT("/orders/:id", or.orderController.UpdateOrder)
	or.router.DELETE("/orders/:id", or.orderController.DeleteOrder)
}
