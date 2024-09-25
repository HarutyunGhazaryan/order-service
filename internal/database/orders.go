package database

import (
	db "OrderService/internal/generated"
	"OrderService/internal/models"
	"context"
	"database/sql"
)

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: true}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

type Queries interface {
	InsertOrder(ctx context.Context, params db.InsertOrderParams) error
	InsertDelivery(ctx context.Context, params db.InsertDeliveryParams) error
	InsertPayment(ctx context.Context, params db.InsertPaymentParams) error
	InsertItem(ctx context.Context, params db.InsertItemParams) error
}

func SaveOrderToDB(ctx context.Context, q Queries, order models.Order) error {
	orderParams := db.InsertOrderParams{
		OrderUid:          order.OrderUID,
		TrackNumber:       ToNullString(order.TrackNumber),
		Entry:             ToNullString(order.Entry),
		Locale:            ToNullString(order.Locale),
		InternalSignature: ToNullString(order.InternalSignature),
		CustomerID:        ToNullString(order.CustomerID),
		DeliveryService:   ToNullString(order.DeliveryService),
		Shardkey:          ToNullString(order.DeliveryService),
		SmID:              sql.NullInt32{Int32: int32(order.SmID), Valid: true},
		DateCreated:       sql.NullTime{Time: order.DateCreated, Valid: true},
		OofShard:          ToNullString(order.OofShard),
	}
	err := q.InsertOrder(ctx, orderParams)

	if err != nil {
		return err
	}

	err = q.InsertDelivery(ctx, db.InsertDeliveryParams{
		OrderUid: ToNullString(order.OrderUID),
		Name:     ToNullString(order.Delivery.Name),
		Phone:    ToNullString(order.Delivery.Phone),
		Zip:      ToNullString(order.Delivery.Zip),
		City:     ToNullString(order.Delivery.City),
		Address:  ToNullString(order.Delivery.Address),
		Region:   ToNullString(order.Delivery.Region),
		Email:    ToNullString(order.Delivery.Email),
	})
	if err != nil {
		return err
	}

	err = q.InsertPayment(ctx, db.InsertPaymentParams{
		OrderUid:     ToNullString(order.OrderUID),
		Transaction:  ToNullString(order.Payment.Transaction),
		RequestID:    ToNullString(order.Payment.RequestID),
		Currency:     ToNullString(order.Payment.Currency),
		Provider:     ToNullString(order.Payment.Provider),
		Amount:       sql.NullInt32{Int32: int32(order.Payment.Amount), Valid: true},
		PaymentDt:    sql.NullInt64{Int64: order.Payment.PaymentDT, Valid: true},
		Bank:         ToNullString(order.Payment.Bank),
		DeliveryCost: sql.NullInt32{Int32: int32(order.Payment.DeliveryCost), Valid: true},
		GoodsTotal:   sql.NullInt32{Int32: int32(order.Payment.GoodsTotal), Valid: true},
		CustomFee:    sql.NullInt32{Int32: int32(order.Payment.CustomFee), Valid: true},
	})
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err = q.InsertItem(ctx, db.InsertItemParams{
			OrderUid:    ToNullString(order.OrderUID),
			ChrtID:      int32(item.ChrtID),
			TrackNumber: ToNullString(item.TrackNumber),
			Price:       sql.NullInt32{Int32: int32(item.Price), Valid: true},
			Rid:         ToNullString(item.Rid),
			Name:        ToNullString(item.Name),
			Sale:        sql.NullInt32{Int32: int32(item.Sale), Valid: true},
			Size:        ToNullString(item.Size),
			TotalPrice:  sql.NullInt32{Int32: int32(item.TotalPrice), Valid: true},
			NmID:        sql.NullInt32{Int32: int32(item.NmID), Valid: true},
			Brand:       ToNullString(item.Brand),
			Status:      sql.NullInt32{Int32: int32(item.Status), Valid: true},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
