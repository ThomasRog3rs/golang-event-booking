package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)

	server.GET("/event/:id", getEventById)
	server.DELETE("/event/:id", deleteEventById)

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

func getEventById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID param could not be parsed"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get item with id: '" + idParam + "'."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": *event})

}

func deleteEventById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID param could not be parsed"})
		return
	}

	err = models.DeleteEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete item with id: '" + idParam + "'."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted item with id: '" + idParam + "'."})
}
