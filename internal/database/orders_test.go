package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	db "OrderService/internal/generated"
	"OrderService/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type MockQueries struct {
	DB *sql.DB
}

func (mq *MockQueries) InsertOrder(ctx context.Context, params db.InsertOrderParams) error {
	_, err := mq.DB.ExecContext(ctx, "INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		params.OrderUid, params.TrackNumber, params.Entry, params.Locale, params.InternalSignature, params.CustomerID, params.DeliveryService, params.Shardkey, params.SmID, params.DateCreated, params.OofShard)
	return err
}

func (mq *MockQueries) InsertDelivery(ctx context.Context, params db.InsertDeliveryParams) error {
	_, err := mq.DB.ExecContext(ctx, "INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		params.OrderUid, params.Name, params.Phone, params.Zip, params.City, params.Address, params.Region, params.Email)
	return err
}

func (mq *MockQueries) InsertPayment(ctx context.Context, params db.InsertPaymentParams) error {
	_, err := mq.DB.ExecContext(ctx, "INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		params.OrderUid, params.Transaction, params.RequestID, params.Currency, params.Provider, params.Amount, params.PaymentDt, params.Bank, params.DeliveryCost, params.GoodsTotal, params.CustomFee)
	return err
}

func (mq *MockQueries) InsertItem(ctx context.Context, params db.InsertItemParams) error {
	_, err := mq.DB.ExecContext(ctx, "INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		params.OrderUid, params.ChrtID, params.TrackNumber, params.Price, params.Rid, params.Name, params.Sale, params.Size, params.TotalPrice, params.NmID, params.Brand, params.Status)
	return err
}

func TestSaveOrderToDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock object: %v", err)
	}
	defer db.Close()

	queries := &MockQueries{db}

	mock.ExpectExec("INSERT INTO orders").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO delivery").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO payment").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO items").WillReturnResult(sqlmock.NewResult(1, 1))

	dateCreated, _ := time.Parse(time.RFC3339, "2021-11-26T06:22:19Z")

	order := models.Order{
		OrderUID:          "b563feb7b2b84b6test",
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       dateCreated,
		OofShard:          "1",
		Delivery: models.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: models.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDT:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []models.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
	}

	err = SaveOrderToDB(context.Background(), queries, order)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Some expectations were not met: %v", err)
	}
}
