package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/resturant/database"
	"github.com/resturant/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	logfile, err := os.Create("./logs/gin.log")
	if err != nil {
		log.Fatal("error creating log file:", err)
	}
	gin.DefaultWriter = io.MultiWriter(logfile)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/home/masoud/resturant-project/gin.log"))
	database.ConnectToDB()
	router.Use(gin.Logger())
	// router.Use(middleware.authentication())
	routes.CustomersRoute(router)
	routes.StaffRoutes(router)
	routes.PaymentRoutes(router)
	routes.OderItemRoutes(router)
	routes.OrderRoutes(router)
	router.Run(":" + port)
}
