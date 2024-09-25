package database

import (
	db "OrderService/internal/generated"
	"context"
)

func DeleteOrderFromDB(ctx context.Context, q db.Queries, orderID string) error {
	return q.DeleteOrder(ctx, orderID)
}
