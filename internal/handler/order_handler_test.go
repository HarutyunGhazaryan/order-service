package handler

import (
	"OrderService/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

type MockOrderStore struct {
	data map[string]models.Order
}

func (m *MockOrderStore) Get(orderUID string) (models.Order, bool) {
	order, found := m.data[orderUID]
	return order, found
}

func (m *MockOrderStore) Add(order models.Order) {
	m.data[order.OrderUID] = order
}

func TestOrderHandler(t *testing.T) {
	mockStore := &MockOrderStore{
		data: map[string]models.Order{
			"123": {
				OrderUID:          "123",
				TrackNumber:       "",
				Entry:             "",
				Delivery:          models.Delivery{},
				Payment:           models.Payment{},
				Items:             nil,
				Locale:            "",
				InternalSignature: "",
				CustomerID:        "",
				DeliveryService:   "",
				Shardkey:          "",
				SmID:              0,
				DateCreated:       time.Time{},
				OofShard:          "",
			},
		},
	}
	handler := OrderHandler(mockStore)

	tests := []struct {
		name           string
		orderUID       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Order",
			orderUID:       "123",
			expectedStatus: http.StatusOK,
			expectedBody: `{
							"order_uid":"123","track_number":"","entry":"",
							"delivery":{"name":"","phone":"","zip":"","city":"","address":"","region":"","email":""},
							"payment":{"transaction":"","request_id":"","currency":"","provider":"","amount":0,"payment_dt":0,"bank":"","delivery_cost":0,"goods_total":0,"custom_fee":0},
							"items":null,"locale":"","internal_signature":"","customer_id":"","delivery_service":"","shardkey":"","sm_id":0,"date_created":"0001-01-01T00:00:00Z","oof_shard":""}
							`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/orders/"+tt.orderUID, nil)
			rec := httptest.NewRecorder()
			router := chi.NewRouter()
			router.Get("/orders/{order_uid}", handler)
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}
