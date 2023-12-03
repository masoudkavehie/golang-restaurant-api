package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/resturant/controllers"
)

func PaymentRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/payments", controllers.GetPayments())
	incomingroutes.GET("/payment/:id", controllers.GetPayment())
	incomingroutes.POST("/payment", controllers.CreatePayment())
	incomingroutes.PUT("/payment/:id", controllers.UpdatePayment())
	incomingroutes.DELETE("/payment/:id", controllers.DeletePayment())
}
