package main

import (
	"OrderService/internal/app"
	"OrderService/internal/config"
	db "OrderService/internal/generated"
	"context"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := app.ConnectToDB(cfg.DB_URL)
	dbQueries := db.New(dbConn)
	orderCache := app.RestoreCache(dbConn)
	kafkaReader := app.ConnectToKafka(cfg.KAFKA_BROKER, cfg.KAFKA_TOPIC)

	app.StartKafkaSubscriber(kafkaReader, orderCache, dbConn)
	ctx := context.Background()
	app.StartHTTPServer(ctx, cfg.PORT, orderCache, dbQueries)
}
