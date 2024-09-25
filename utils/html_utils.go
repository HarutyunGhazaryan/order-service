package utils

import (
	"OrderService/internal/models"
	"fmt"
	"strings"
	"time"
)

func GenerateTable(title string, data map[string]interface{}, keys []string) string {
	var sb strings.Builder
	sb.WriteString(`<h2>` + title + `</h2>
    <table>
      <tbody>`)
	for _, k := range keys {
		sb.WriteString(`<tr><td>` + k + `</td><td>` + valueToString(data[k]) + `</td></tr>`)
	}
	sb.WriteString(`</tbody>
    </table>`)
	return sb.String()
}

func GenerateItemsTables(items []models.Item) string {
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(`
    <table>
      <tbody>`)
		itemData := map[string]interface{}{
			"ChrtID":      item.ChrtID,
			"TrackNumber": item.TrackNumber,
			"Price":       item.Price,
			"Rid":         item.Rid,
			"Name":        item.Name,
			"Sale":        item.Sale,
			"Size":        item.Size,
			"TotalPrice":  item.TotalPrice,
			"NmID":        item.NmID,
			"Brand":       item.Brand,
			"Status":      item.Status,
		}
		keys := []string{"ChrtID", "TrackNumber", "Price", "Rid", "Name", "Sale", "Size", "TotalPrice", "NmID", "Brand", "Status"}
		for _, k := range keys {
			sb.WriteString(`<tr><td>` + k + `</td><td>` + valueToString(itemData[k]) + `</td></tr>`)
		}
		sb.WriteString(`</tbody>
    </table>`)
	}
	return sb.String()
}

func valueToString(value interface{}) string {
	switch v := value.(type) {
	case float64:
		return fmt.Sprintf("%.2f", v)
	case string:
		return v
	case time.Time:
		return v.Format(time.RFC3339)
	default:
		return fmt.Sprintf("%v", v)
	}
}
