package main

import (
	"net/http"

	model "github.com/Mautu/altermangetokafka/modle"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/altermange", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": " successful get altermange message"})

	})
	router.Run()
}
