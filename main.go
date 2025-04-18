package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8000")

}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"message": "could not parse request data"})
		return
	}

	event.UserID = 1
	event.ID = 1

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event create successfully", "data": event})
}
