package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jrgmonsalve/back-event-booking/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authorized := server.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.POST("/events", createEvent)
	authorized.PUT("/events/:id", updateEvent)
	authorized.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
