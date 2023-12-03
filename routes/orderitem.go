package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resturant/controllers"
)

func OderItemRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/orderitems", controllers.GetOrderItems)
	incomingroutes.GET("/orderitem/:id", controllers.GetOrderItem)
	incomingroutes.POST("/orderitems", controllers.CreateOrderItem)
	incomingroutes.DELETE("/oderitem/:id", controllers.DeletedOrderItem)
}
