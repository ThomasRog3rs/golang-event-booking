package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)

	server.Run(":8080")
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	//dummy ids for now
	event.ID = 1
	event.UserID = 1

	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get events from datbase"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": events})
}
