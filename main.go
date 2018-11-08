package main

import (
	"memento/pkg/util"
	"net/http"
	"strconv"
	"time"

	"github.com/Mautu/altermangetokafka/config"
	"github.com/Mautu/altermangetokafka/sender"
	"go.uber.org/zap"

	model "github.com/Mautu/altermangetokafka/modle"
	"github.com/gin-gonic/gin"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	router := gin.Default()
	router.POST("/altermange", func(c *gin.Context) {
		var notification model.Notification
		err := c.BindJSON(&notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logger.Info("kafka config load successful",
			zap.String("host", config.Kafka.Host[0]+config.Kafka.Host[1]+config.Kafka.Host[2]),
			zap.String("topic", config.Kafka.Topic),
			zap.Duration("backoff", time.Second),
		)
		sender.Sender(config.Kafka.Host, notification, config.Kafka.Topic)
		c.JSON(http.StatusOK, gin.H{"message": " successful get altermange message"})
	})
	router.Run(util.StrAddbybyte(":", strconv.Itoa(config.Port)))
}
