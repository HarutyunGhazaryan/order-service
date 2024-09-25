package cache

import (
	"OrderService/utils"
	"database/sql"
)

func (c *OrderCache) RestoreFromDB(db *sql.DB) error {
	rows, err := utils.QueryOrders(db)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		order, err := utils.ScanOrder(rows)
		if err != nil {
			return err
		}

		c.Add(order)

		if err := c.loadRelatedData(db, order.OrderUID); err != nil {
			return err
		}
	}
	return nil
}

func (c *OrderCache) loadRelatedData(db *sql.DB, orderUID string) error {
	delivery, err := utils.ScanDelivery(db, orderUID)
	if err != nil {
		return err
	}

	payment, err := utils.ScanPayment(db, orderUID)
	if err != nil {
		return err
	}

	items, err := utils.ScanItems(db, orderUID)
	if err != nil {
		return err
	}

	order := c.orders[orderUID]
	order.Delivery = delivery
	order.Payment = payment
	order.Items = items
	c.Add(order)

	return nil
}
