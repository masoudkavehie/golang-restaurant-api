package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/resturant/database"
	"github.com/resturant/models"
)

func GetOrderItems(c *gin.Context) {
	var allorderitems []models.OrderItem
	database.DB.Find(&allorderitems)
	c.JSON(200, gin.H{
		"allorderitems": allorderitems,
	})
}

func GetOrderItem(c *gin.Context) {
	id := c.Param("id")
	var targetitem models.OrderItem
	database.DB.First(&targetitem, id)
	c.JSON(200, gin.H{
		"targetitem": targetitem,
	})

}
func CreateOrderItem(c *gin.Context) {
	var orderitemcreate models.OrderItem
	c.Bind(&orderitemcreate)
	create := models.OrderItem{OrderID: orderitemcreate.OrderID, ItemID: orderitemcreate.ItemID, Quantity: orderitemcreate.Quantity, Subtotal: orderitemcreate.Subtotal}
	result := database.DB.Create(&create)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"created": create,
	})

}
func UpdateOrderitem(c *gin.Context) {
	var orderitem models.OrderItem
	database.DB.Where("id=?", c.Param("id")).First(&orderitem)
	c.BindJSON(&orderitem)
	database.DB.Save(&orderitem)
	c.JSON(200, &orderitem)
}
func DeletedOrderItem(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.OrderItem{}, id)
	c.Status(200)
}

// func GetOrderItemById() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }
