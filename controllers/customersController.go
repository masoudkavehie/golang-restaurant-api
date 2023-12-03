package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/resturant/database"
	"github.com/resturant/models"
)

func GetCustomers(c *gin.Context) {
	var Customers []models.Customer
	database.DB.Find(&Customers)
	c.JSON(200, gin.H{
		"Customers": Customers,
	})
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var Customer []models.Customer
	database.DB.Find(&Customer, id)
	c.JSON(200, gin.H{
		"Customer": Customer,
	})

}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	database.DB.Create(&customer)
	c.JSON(200, gin.H{
		"payment": customer,
	})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	database.DB.Where("id=?", c.Param("id")).First(&customer)
	c.BindJSON(&customer)
	database.DB.Save(&customer)
	c.JSON(200, &customer)
}
func DeletedCusomer(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Customer{}, id)
	c.Status(200)
}
