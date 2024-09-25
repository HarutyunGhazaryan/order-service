package app

import (
	"OrderService/internal/cache"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
)

func ConnectToDB(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	return db
}

func RestoreCache(db *sql.DB) *cache.OrderCache {
	orderCache := cache.NewOrderCache()
	if err := orderCache.RestoreFromDB(db); err != nil {
		log.Fatalf("Failed to restore cache from database: %v", err)
	}
	return orderCache
}

func ConnectToKafka(broker, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "order-service-group",
	})
}
