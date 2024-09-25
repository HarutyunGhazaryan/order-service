package handler

import (
	"OrderService/internal/cache"
	"OrderService/internal/database"
	db "OrderService/internal/generated"
	"OrderService/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func HandleKafkaMessage(msg kafka.Message, c *cache.OrderCache, dbConn *sql.DB) {
	var order models.Order
	if err := json.Unmarshal(msg.Value, &order); err != nil {
		log.Printf("Failed to unmarshal order: %v", err)
		return
	}

	ctx := context.Background()

	queries := db.New(dbConn)

	if err := database.SaveOrderToDB(ctx, queries, order); err != nil {
		log.Printf("Failed to save order to DB: %v", err)
		return
	}

	c.Add(order)
}
