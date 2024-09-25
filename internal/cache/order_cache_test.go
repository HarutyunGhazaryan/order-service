package cache

import (
	"OrderService/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderCacheAddAndGet(t *testing.T) {
	cache := NewOrderCache()

	order := models.Order{
		OrderUID: "order1",
	}
	cache.Add(order)

	retrievedOrder, found := cache.Get("order1")
	assert.True(t, found)
	assert.Equal(t, order, retrievedOrder)
}

func TestOrderCacheGetNonExistentOrder(t *testing.T) {
	cache := NewOrderCache()

	_, found := cache.Get("order2")
	assert.False(t, found)
}
