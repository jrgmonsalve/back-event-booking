package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrgmonsalve/back-event-booking/db"
	"github.com/jrgmonsalve/back-event-booking/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	err = event.Save()
	if err != nil {
		log.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		c.JSON(http.StatusInternalServerError, gin.H{"error---------->>>>>": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}
