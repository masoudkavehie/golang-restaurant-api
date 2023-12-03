package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/resturant/controllers"
)

func StaffRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/staffs", controllers.GetStaffs())
	incomingroutes.GET("/staff/:id", controllers.GetStaff())
	incomingroutes.POST("/staff", controllers.CreateStaff())
	incomingroutes.PUT("/staff/:id", controllers.UpdateStaff())
}
