package cache

import (
	"OrderService/internal/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRestoreFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred while opening a mock database connection", err)
	}
	defer db.Close()

	orderCache := &OrderCache{orders: make(map[string]models.Order)}

	dateCreated, _ := time.Parse("2006-01-02", "2024-09-04")
	paymentDate := dateCreated.Unix()

	rows := sqlmock.NewRows([]string{
		"order_uid", "track_number", "entry", "locale", "internal_signature",
		"customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard",
	}).AddRow("order1", "track1", "entry1", "locale1", "signature1", "customer1", "service1", "shard1", 1, dateCreated, "oof1")

	mock.ExpectQuery(`^SELECT order_uid, track_number, entry, locale, COALESCE\(internal_signature, ''\), customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders$`).WillReturnRows(rows)

	mock.ExpectQuery(`^SELECT name, phone, zip, city, address, region, email FROM delivery WHERE order_uid = \$1$`).WithArgs("order1").WillReturnRows(sqlmock.NewRows([]string{
		"name", "phone", "zip", "city", "address", "region", "email",
	}).AddRow("name1", "phone1", "zip1", "city1", "address1", "region1", "email1"))

	mock.ExpectQuery(`^SELECT transaction, COALESCE\(request_id, ''\), currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment WHERE order_uid = \$1$`).WithArgs("order1").WillReturnRows(sqlmock.NewRows([]string{
		"transaction", "request_id", "currency", "provider", "amount", "payment_dt", "bank", "delivery_cost", "goods_total", "custom_fee",
	}).AddRow("transaction1", "request1", "USD", "provider1", 100.0, paymentDate, "bank1", 10.0, 90.0, 5.0))

	mock.ExpectQuery(`^SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = \$1$`).WithArgs("order1").WillReturnRows(sqlmock.NewRows([]string{
		"chrt_id", "track_number", "price", "rid", "name", "sale", "size", "total_price", "nm_id", "brand", "status",
	}).AddRow(12, "track1", 50.0, "rid1", "item1", 5, "size1", 250.0, 12, "brand1", 1))

	if err := orderCache.RestoreFromDB(db); err != nil {
		t.Fatalf("an error '%s' occurred while restoring data from the database", err)
	}

	order, exists := orderCache.orders["order1"]
	if !exists {
		t.Errorf("expected order with UID 'order1' to exist")
		return
	}

	assert.Equal(t, "order1", order.OrderUID)
	assert.Equal(t, "track1", order.TrackNumber)
	assert.Equal(t, "name1", order.Delivery.Name)
	assert.Len(t, order.Items, 1)
}
