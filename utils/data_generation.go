package utils

import (
	"OrderService/internal/models"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomOrder() models.Order {
	return models.Order{
		OrderUID:    uuid.New().String(),
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: models.Delivery{
			Name:    generateRandomName(),
			Phone:   generateRandomPhoneNumber(),
			Zip:     generateRandomZip(),
			City:    generateRandomCity(),
			Address: generateRandomAddress(),
			Region:  generateRandomRegion(),
			Email:   generateRandomEmail(),
		},
		Payment: models.Payment{
			Transaction:  uuid.New().String(),
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Intn(5000) + 100,
			PaymentDT:    time.Now().Unix(),
			Bank:         "alpha",
			DeliveryCost: rand.Intn(2000),
			GoodsTotal:   rand.Intn(500),
			CustomFee:    0,
		},
		Items: []models.Item{
			{
				ChrtID:      rand.Intn(9999999),
				TrackNumber: "WBILMTESTTRACK",
				Price:       rand.Intn(500) + 1,
				Rid:         uuid.New().String(),
				Name:        generateRandomString(10),
				Sale:        rand.Intn(100),
				Size:        "0",
				TotalPrice:  rand.Intn(500),
				NmID:        rand.Intn(9999999),
				Brand:       generateRandomString(5),
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
}
