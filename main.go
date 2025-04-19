package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/database"
	"github.com/sanoy24/event-booking-rest-api/routes"
)

func main() {
	databse.InitializeDb()
	server := gin.Default()
	routes.InitRoutes(server)

	server.Run(":8000")

}
