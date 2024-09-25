package main

import (
	"OrderService/internal/config"
	"OrderService/utils"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	cfg := config.LoadConfig()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.KAFKA_BROKER},
		Topic:    cfg.KAFKA_TOPIC,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	order := utils.GenerateRandomOrder()
	data, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	_, err = utils.ValidateOrderData(data)
	if err != nil {
		log.Fatalf("Order validation failed: %v", err)
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Order published successfully")
}
