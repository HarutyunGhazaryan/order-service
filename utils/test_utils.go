package utils

import (
	"OrderService/internal/database"
	db "OrderService/internal/generated"
	"OrderService/internal/models"
	"context"
	"fmt"
)

func CreateTestOrder(ctx context.Context, queries *db.Queries) (models.Order, error) {
	order := GenerateRandomOrder()
	err := database.SaveOrderToDB(ctx, queries, order)
	if err != nil {
		fmt.Printf("Error saving order to DB: %v\n", err)
		return models.Order{}, err
	}
	return order, err
}

func DeleteTestOrder(ctx context.Context, queries *db.Queries, orderUID string) error {
	err := database.DeleteOrderFromDB(ctx, *queries, orderUID)
	if err != nil {
		fmt.Printf("Error deleting order form DB: %V\n", err)
		return err
	}
	return nil
}
