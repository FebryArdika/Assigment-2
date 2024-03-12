package routers

import (
	ctrl "Assignment-2-Golang/controllers"

	"github.com/gin-gonic/gin"
)

func ServerOn() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", ctrl.InsertOrders)
	router.GET("/orders/", ctrl.ShowOrders)
	router.PUT("/orders/:OrderID", ctrl.UpdateOrder)
	router.DELETE("/orders/:OrderID", ctrl.DeleteOrder)
	return router
}
