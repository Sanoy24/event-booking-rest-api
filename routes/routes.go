package routes

import "github.com/gin-gonic/gin"

func eventsRoutes(eventGroup *gin.RouterGroup) {

	events := eventGroup.Group("/events")
	{
		events.GET("/", getEvents)
		events.GET("/:id", getSingleEvent)
		events.POST("/", createEvent)
		events.PATCH("/:id", updateEvent)
		events.DELETE("/:id", deleteEvent)
	}
}
