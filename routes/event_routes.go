package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/handlers"
)

func eventsRoutes(eventGroup *gin.RouterGroup) {

	events := eventGroup.Group("/events")
	{
		events.GET("/", handlers.GetEvents)
		events.GET("/:id", handlers.GetSingleEvent)
		events.POST("/", handlers.CreateEvent)
		events.PATCH("/:id", handlers.UpdateEvent)
		events.DELETE("/:id", handlers.DeleteEvent)
	}
}
