package handler

import (
	db "OrderService/internal/generated"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func DBOrderHandler(dbQueries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderUID := chi.URLParam(r, "order_uid")

		order, err := dbQueries.GetOrderByUID(r.Context(), orderUID)
		if err != nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}
