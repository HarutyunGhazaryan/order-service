package utils

import (
	"OrderService/internal/models"
	"time"
)

func ConvertOrderToMap(order models.Order) map[string]interface{} {
	return map[string]interface{}{
		"OrderUID":        order.OrderUID,
		"TrackNumber":     order.TrackNumber,
		"Entry":           order.Entry,
		"Locale":          order.Locale,
		"CustomerID":      order.CustomerID,
		"DeliveryService": order.DeliveryService,
		"Shardkey":        order.Shardkey,
		"SmID":            order.SmID,
		"DateCreated":     order.DateCreated.Format(time.RFC3339),
		"OofShard":        order.OofShard,
	}
}

func ConvertDeliveryToMap(delivery models.Delivery) map[string]interface{} {
	return map[string]interface{}{
		"Name":    delivery.Name,
		"Phone":   delivery.Phone,
		"Zip":     delivery.Zip,
		"City":    delivery.City,
		"Address": delivery.Address,
		"Region":  delivery.Region,
		"Email":   delivery.Email,
	}
}

func ConvertPaymentToMap(payment models.Payment) map[string]interface{} {
	return map[string]interface{}{
		"Transaction":  payment.Transaction,
		"RequestID":    payment.RequestID,
		"Currency":     payment.Currency,
		"Provider":     payment.Provider,
		"Amount":       payment.Amount,
		"PaymentDT":    time.Unix(payment.PaymentDT, 0).Format(time.RFC3339),
		"Bank":         payment.Bank,
		"DeliveryCost": payment.DeliveryCost,
		"GoodsTotal":   payment.GoodsTotal,
		"CustomFee":    payment.CustomFee,
	}
}
