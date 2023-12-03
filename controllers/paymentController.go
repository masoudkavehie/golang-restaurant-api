package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/resturant/database"
	"github.com/resturant/models"
)

func GetPayments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payments []models.Payment
		database.DB.Find(&payments)
		c.JSON(200, gin.H{
			"payments": &payments,
		})
	}
}

func GetPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payment models.Payment
		database.DB.First(&payment, id)
		c.JSON(200, gin.H{
			"payment": payment,
		})
	}
}

func UpdatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payment models.Payment

		// Find the payment record by ID
		result := database.DB.Where("\"paymentID\" = ?", id).First(&payment)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Payment not found"})
			return
		}

		// Bind JSON data to the payment model
		c.BindJSON(&payment)

		// Save the updated payment record
		database.DB.Save(&payment)

		// Return the updated payment as JSON response
		c.JSON(200, gin.H{
			"payment": payment,
		})
	}
}

func CreatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payment models.Payment
		c.BindJSON(&payment)
		database.DB.Create(&payment)
		c.JSON(200, gin.H{
			"payment": payment,
		})
	}

}
func DeletePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result := database.DB.Delete(&models.Payment{}, id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.Status(200)
	}
}
