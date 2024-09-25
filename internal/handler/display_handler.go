package handler

import (
	"OrderService/internal/cache"
	"OrderService/utils"
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func DisplayHandler(c cache.OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderUID := chi.URLParam(r, "order_uid")

		tmpl, err := template.ParseFiles("internal/templates/display.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		data := struct {
			OrderHTML template.HTML
			Error     string
		}{}

		if orderUID != "" {
			order, found := c.Get(orderUID)
			if !found {
				data.Error = "Order not found"
			} else {
				var sb strings.Builder
				sb.WriteString(`<div class="tables-container">`)
				sb.WriteString(`<div class="table-wrapper">`)
				sb.WriteString(utils.GenerateTable("Order", utils.ConvertOrderToMap(order), []string{
					"OrderUID", "TrackNumber", "Entry", "Locale", "CustomerID", "DeliveryService", "Shardkey", "SmID", "DateCreated", "OofShard"}))
				sb.WriteString(`</div>`)
				sb.WriteString(`<div class="table-wrapper">`)
				sb.WriteString(utils.GenerateTable("Delivery", utils.ConvertDeliveryToMap(order.Delivery), []string{
					"Name", "Phone", "Zip", "City", "Address", "Region", "Email"}))
				sb.WriteString(`</div>`)
				sb.WriteString(`</div>`)

				sb.WriteString(`<div class="tables-container">`)
				sb.WriteString(`<div class="table-wrapper">`)
				sb.WriteString(utils.GenerateTable("Payment", utils.ConvertPaymentToMap(order.Payment), []string{
					"Transaction", "RequestID", "Currency", "Provider", "Amount", "PaymentDT", "Bank", "DeliveryCost", "GoodsTotal", "CustomFee"}))
				sb.WriteString(`</div>`)
				sb.WriteString(`<div class="table-wrapper">`)
				sb.WriteString(`<h2>Items</h2>`)
				sb.WriteString(utils.GenerateItemsTables(order.Items))
				sb.WriteString(`</div>`)
				sb.WriteString(`</div>`)

				data.OrderHTML = template.HTML(sb.String())
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}
