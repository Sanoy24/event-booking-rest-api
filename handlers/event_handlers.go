package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/models"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
	}
	ctx.JSON(http.StatusOK, events)
}

func GetSingleEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invaid id"})
		return
	}

	events, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "errot occured while fetching the data, with err", "error": err})
		return
	}

	res := Response{
		Status:  "success",
		Message: "event fetched successfully",
		Data:    events,
	}
	ctx.JSON(http.StatusOK, res)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "could not parse request data"})
		return
	}
	userId := ctx.GetInt64("userId")
	event.UserID = int(userId)

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event create successfully", "data": event})
}

func UpdateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invaid id"})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "errot occured while fetching the data, with err", "error": err})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {

		ctx.JSON(http.StatusOK, gin.H{"message": "could not parse request data"})
		return
	}

	userId := ctx.GetInt64("userId")

	if userId != int64(updatedEvent.UserID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to update event"})
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not update the request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated event successfully"})
}

func DeleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invaid id"})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "errot occured while fetching the data, with err", "error": err})
		return
	}
	var event models.Event
	event.ID = id

	userId := ctx.GetInt64("userId")

	if userId != int64(event.UserID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to delete event"})
	}

	err = event.DeleteEvent()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can not delete event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})

}
