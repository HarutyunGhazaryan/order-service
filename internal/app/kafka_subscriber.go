package app

import (
	"OrderService/internal/cache"
	"OrderService/internal/handler"
	"context"
	"database/sql"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartKafkaSubscriber(reader *kafka.Reader, orderCache *cache.OrderCache, db *sql.DB) {
	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message from Kafka: %v", err)
				continue
			}

			handler.HandleKafkaMessage(msg, orderCache, db)
		}
	}()
}
