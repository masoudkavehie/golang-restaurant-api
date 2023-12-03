package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/resturant/database"
	"github.com/resturant/models"
)

func GetOders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Orders []models.Order
		database.DB.Find(&Orders)
		c.JSON(200, gin.H{
			"Order:": Orders,
		})
	}
}

func GetOder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order models.Order
		database.DB.First(&order, id)
		c.JSON(200, gin.H{
			"order:": order,
		})

	}
}

func UpdateOder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order models.Order
		var orders models.Order
		c.Bind(&order)
		database.DB.First(&orders, id)
		database.DB.Model(&orders).Updates(&order)
		c.JSON(200, gin.H{
			"Order": order,
		})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		c.Bind(&order)
		database.DB.Create(&order)
		c.JSON(200, gin.H{
			"order": order,
		})
	}
}

func DeletedOrder(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Order{}, id)
	c.Status(200)
}
