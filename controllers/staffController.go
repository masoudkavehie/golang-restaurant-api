package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/resturant/database"
	"github.com/resturant/models"
)

func GetStaffs() gin.HandlerFunc {
	return func(c *gin.Context) {
		var staffs []models.Staff
		database.DB.Find(&staffs)
		c.JSON(200, gin.H{
			"Customers": staffs,
		})
	}
}

func GetStaff() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var staff []models.Staff
		database.DB.Find(&staff, id)
		c.JSON(200, gin.H{
			"Customer": staff,
		})
	}
}

func UpdateStaff() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateStaff() gin.HandlerFunc {
	return func(c *gin.Context) {
		var staff models.Staff
		c.BindJSON(&staff)
		database.DB.Create(&staff)
		c.JSON(200, gin.H{
			"payment": staff,
		})
	}
}
func DeleteStaff() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		database.DB.Delete(&models.Staff{}, id)
		c.Status(200)
	}
}
