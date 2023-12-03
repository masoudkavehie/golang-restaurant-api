package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/resturant/controllers"
)

func CustomersRoute(incomingroutes *gin.Engine) {
	incomingroutes.GET("/customers", controllers.GetCustomers)
	incomingroutes.GET("/customers/:id", controllers.GetCustomer)
	incomingroutes.POST("/customers", controllers.CreateCustomer)
	incomingroutes.PUT("/customer/:id", controllers.CreateCustomer)

	incomingroutes.DELETE("/customers/:id", controllers.DeletedCusomer)

}
