package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // GET,POST, PUT,PATCH,DELETE
	server.GET("/events/:id", getEvent) // /events/1 //evetns/23

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	//this methods runs left to right. so that there is a better way when we dealing with more than one
	// protected routes. So that we can see the ussage of the on the upper part

	server.POST("/signup", signup)
	server.POST("/login", login)

}
