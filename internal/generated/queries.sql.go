// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
)

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders WHERE order_uid = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderUid string) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, orderUid)
	return err
}

const getAllOrders = `-- name: GetAllOrders :many
SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders
`

func (q *Queries) GetAllOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getAllOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderUid,
			&i.TrackNumber,
			&i.Entry,
			&i.Locale,
			&i.InternalSignature,
			&i.CustomerID,
			&i.DeliveryService,
			&i.Shardkey,
			&i.SmID,
			&i.DateCreated,
			&i.OofShard,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderByUID = `-- name: GetOrderByUID :one
SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders WHERE order_uid = $1
`

func (q *Queries) GetOrderByUID(ctx context.Context, orderUid string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByUID, orderUid)
	var i Order
	err := row.Scan(
		&i.OrderUid,
		&i.TrackNumber,
		&i.Entry,
		&i.Locale,
		&i.InternalSignature,
		&i.CustomerID,
		&i.DeliveryService,
		&i.Shardkey,
		&i.SmID,
		&i.DateCreated,
		&i.OofShard,
	)
	return i, err
}

const insertDelivery = `-- name: InsertDelivery :exec
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
`

type InsertDeliveryParams struct {
	OrderUid sql.NullString
	Name     sql.NullString
	Phone    sql.NullString
	Zip      sql.NullString
	City     sql.NullString
	Address  sql.NullString
	Region   sql.NullString
	Email    sql.NullString
}

func (q *Queries) InsertDelivery(ctx context.Context, arg InsertDeliveryParams) error {
	_, err := q.db.ExecContext(ctx, insertDelivery,
		arg.OrderUid,
		arg.Name,
		arg.Phone,
		arg.Zip,
		arg.City,
		arg.Address,
		arg.Region,
		arg.Email,
	)
	return err
}

const insertItem = `-- name: InsertItem :exec
INSERT INTO items (
    order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
`

type InsertItemParams struct {
	OrderUid    sql.NullString
	ChrtID      int32
	TrackNumber sql.NullString
	Price       sql.NullInt32
	Rid         sql.NullString
	Name        sql.NullString
	Sale        sql.NullInt32
	Size        sql.NullString
	TotalPrice  sql.NullInt32
	NmID        sql.NullInt32
	Brand       sql.NullString
	Status      sql.NullInt32
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) error {
	_, err := q.db.ExecContext(ctx, insertItem,
		arg.OrderUid,
		arg.ChrtID,
		arg.TrackNumber,
		arg.Price,
		arg.Rid,
		arg.Name,
		arg.Sale,
		arg.Size,
		arg.TotalPrice,
		arg.NmID,
		arg.Brand,
		arg.Status,
	)
	return err
}

const insertOrder = `-- name: InsertOrder :exec
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature,
    customer_id, delivery_service, shardkey, sm_id, date_created,
    oof_shard
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
`

type InsertOrderParams struct {
	OrderUid          string
	TrackNumber       sql.NullString
	Entry             sql.NullString
	Locale            sql.NullString
	InternalSignature sql.NullString
	CustomerID        sql.NullString
	DeliveryService   sql.NullString
	Shardkey          sql.NullString
	SmID              sql.NullInt32
	DateCreated       sql.NullTime
	OofShard          sql.NullString
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) error {
	_, err := q.db.ExecContext(ctx, insertOrder,
		arg.OrderUid,
		arg.TrackNumber,
		arg.Entry,
		arg.Locale,
		arg.InternalSignature,
		arg.CustomerID,
		arg.DeliveryService,
		arg.Shardkey,
		arg.SmID,
		arg.DateCreated,
		arg.OofShard,
	)
	return err
}

const insertPayment = `-- name: InsertPayment :exec
INSERT INTO payment (
    order_uid, transaction, request_id, currency, provider, amount, payment_dt,
    bank, delivery_cost, goods_total, custom_fee
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
`

type InsertPaymentParams struct {
	OrderUid     sql.NullString
	Transaction  sql.NullString
	RequestID    sql.NullString
	Currency     sql.NullString
	Provider     sql.NullString
	Amount       sql.NullInt32
	PaymentDt    sql.NullInt64
	Bank         sql.NullString
	DeliveryCost sql.NullInt32
	GoodsTotal   sql.NullInt32
	CustomFee    sql.NullInt32
}

func (q *Queries) InsertPayment(ctx context.Context, arg InsertPaymentParams) error {
	_, err := q.db.ExecContext(ctx, insertPayment,
		arg.OrderUid,
		arg.Transaction,
		arg.RequestID,
		arg.Currency,
		arg.Provider,
		arg.Amount,
		arg.PaymentDt,
		arg.Bank,
		arg.DeliveryCost,
		arg.GoodsTotal,
		arg.CustomFee,
	)
	return err
}
