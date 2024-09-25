package utils

import (
	"OrderService/internal/models"
	"database/sql"
)

func QueryOrders(db *sql.DB) (*sql.Rows, error) {
	query := `
		SELECT 
			order_uid, track_number, entry, locale, 
			COALESCE(internal_signature, ''), customer_id, 
			delivery_service, shardkey, sm_id, 
			date_created, oof_shard
		FROM orders
	`
	return db.Query(query)
}

func ScanOrder(rows *sql.Rows) (models.Order, error) {
	var order models.Order
	if err := rows.Scan(
		&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale,
		&order.InternalSignature, &order.CustomerID, &order.DeliveryService,
		&order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard,
	); err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func ScanDelivery(db *sql.DB, orderUID string) (models.Delivery, error) {
	var delivery models.Delivery
	query := `
		SELECT 
			name, phone, zip, city, address, region, email
		FROM delivery WHERE order_uid = $1
	`

	err := db.QueryRow(query, orderUID).Scan(
		&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City,
		&delivery.Address, &delivery.Region, &delivery.Email,
	)
	if err != nil {
		return models.Delivery{}, err
	}
	return delivery, nil
}

func ScanPayment(db *sql.DB, orderUID string) (models.Payment, error) {
	var payment models.Payment
	query := `
		SELECT 
			transaction, COALESCE(request_id, ''), currency, provider, amount, 
			payment_dt, bank, delivery_cost, goods_total, custom_fee
		FROM payment WHERE order_uid = $1
	`

	err := db.QueryRow(query, orderUID).Scan(
		&payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider,
		&payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost,
		&payment.GoodsTotal, &payment.CustomFee,
	)
	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func ScanItems(db *sql.DB, orderUID string) ([]models.Item, error) {
	query := `
		SELECT 
			chrt_id, track_number, price, rid, name, sale, 
			size, total_price, nm_id, brand, status
		FROM items WHERE order_uid = $1
	`
	rows, err := db.Query(query, orderUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(
			&item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name,
			&item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand,
			&item.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
