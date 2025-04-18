package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/database"
	"github.com/sanoy24/event-booking-rest-api/models"
)

func main() {
	databse.InitializeDb()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8000")

}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
	}
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

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event create successfully", "data": event})
}
