package cache

import (
	"OrderService/internal/models"
	"sync"
)

type OrderStore interface {
	Add(order models.Order)
	Get(orderUID string) (models.Order, bool)
}

type OrderCache struct {
	mu     sync.RWMutex
	orders map[string]models.Order
}

func NewOrderCache() *OrderCache {
	return &OrderCache{
		orders: make(map[string]models.Order),
	}
}

func (c *OrderCache) Add(order models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *OrderCache) Get(orderUID string) (models.Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, found := c.orders[orderUID]
	return order, found
}
