package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/resturant/controllers"
)

func OrderRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/orders", controllers.GetOders())
	incomingroutes.GET("/order/:id", controllers.GetOder())
	incomingroutes.POST("/orders", controllers.CreateOrder())
	incomingroutes.PUT("/order/:id", controllers.UpdateOder())
	incomingroutes.DELETE("/order/:id", controllers.DeletedOrder)
}
