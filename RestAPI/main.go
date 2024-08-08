package main

import (
	"net/http"
	"restAPI/db"
	"restAPI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)     // If we have a GET to the /events page, getEvents func is call
	server.GET("/events/:id", getEvent)  //events/1, events/7...
	server.POST("/events", createEvents) // If we have a POST to the /events page, createEvents func is call
	server.Run(":8080")                  //localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, events) // Send Json response to the client
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id"})
		return
	}
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // put the receive data from the POST to event
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
