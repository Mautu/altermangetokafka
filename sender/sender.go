package sender

import (
	"encoding/json"
	"strconv"
	"time"

	log "github.com/Mautu/altermangetokafka/logger"
	modle "github.com/Mautu/altermangetokafka/modle"
	"github.com/Mautu/altermangetokafka/transformer"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func Sender(address []string, notification modle.Notification, topic string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Logger.Error("connect kafka error arama.NewSyncProducer",
			// Structured context as strongly typed Field values.
			zap.Error(err))
		return
	}
	defer p.Close()
	for _, alert := range notification.Alerts {
		messgae, err := json.Marshal(transformer.AlertToMessagekafka(alert))
		if err != nil {
			log.Logger.Error("AlertToMessagekafka error",
				// Structured context as strongly typed Field values.
				zap.Error(err))
		}
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(messgae),
		}
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Logger.Error("send message error",
				// Structured context as strongly typed Field values.
				zap.String("messgae", string(messgae)),
				zap.Error(err),
				zap.Duration("backoff", time.Second))
		} else {
			log.Logger.Info("send message successful",
				// Structured context as strongly typed Field values.
				zap.String("messgae", string(messgae)),
				zap.String("part", strconv.FormatInt(int64(part), 10)),
				zap.Int64("offset", offset))
		}
		time.Sleep(2 * time.Second)
	}
}
