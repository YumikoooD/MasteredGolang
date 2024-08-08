package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)     // If we have a GET to the /events page, getEvents func is call
	server.GET("/events/:id", getEvent)  //events/1, events/7...
	server.POST("/events", createEvents) // If we have a POST to the /events page, createEvents func is call
	server.PUT("/events/:id", updateEvents)
	server.DELETE("/events/:id", deleteEvents)
	server.POST("/signup", signup) // create user
	server.POST("/login", login)   // create user
}
