package utils

import (
	"OrderService/internal/models"
	"encoding/json"
	"fmt"
)

func ValidateOrderData(data []byte) (*models.Order, error) {
	var order models.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return nil, err
	}
	if order.OrderUID == "" {
		return nil, fmt.Errorf("invalid order data: missing order_uid")
	}
	if order.Delivery == (models.Delivery{}) {
		return nil, fmt.Errorf("invalid order data: missing delivery information")
	}
	if order.Payment == (models.Payment{}) {
		return nil, fmt.Errorf("invalid order data: missing payment information")
	}
	if len(order.Items) == 0 {
		return nil, fmt.Errorf("invalid order data: missing items")
	}
	for _, item := range order.Items {
		if item.ChrtID == 0 {
			return nil, fmt.Errorf("invalid order data: item missing ChrtID")
		}
		if item.TrackNumber == "" {
			return nil, fmt.Errorf("invalid order data: item missing TrackNumber")
		}
		if item.Price <= 0 {
			return nil, fmt.Errorf("invalid order data: item price must be positive")
		}
	}

	return &order, nil
}
