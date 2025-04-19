package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanoy24/event-booking-rest-api/handlers"
	"github.com/sanoy24/event-booking-rest-api/middleware"
)

func eventsRoutes(eventGroup *gin.RouterGroup) {
	events := eventGroup.Group("/events")

	events.GET("/", handlers.GetEvents)
	events.GET("/:id", handlers.GetSingleEvent)

	protected := events.Group("/")
	protected.Use(middleware.Authenticate)
	{
		protected.POST("", handlers.CreateEvent)
		protected.PATCH("/:id", handlers.UpdateEvent)
		protected.DELETE("/:id", handlers.DeleteEvent)
	}

}
