package cache

import (
	"sync"
	"time"
)

// Item represents a cached response along with its expiration time
type Item struct {
	Value      []byte
	Expiration int64
}

// MemoryCache represents a thread-safe in-memory cache
type MemoryCache struct {
	items map[string]Item
	mu    sync.RWMutex
}

// NewMemoryCache initializes a new MemoryCache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		items: make(map[string]Item),
	}
}

// Set adds a new item to the cache with the given TTL
func (c *MemoryCache) Set(key string, value []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}

// Get retrieves an item from the cache. Returns nil if not found or expired
func (c *MemoryCache) Get(key string) []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil
	}

	if time.Now().UnixNano() > item.Expiration {
		// Avoid blocking readers to delete, rely on Set to overwrite or occasional cleanup
		return nil
	}

	return item.Value
}

// Clear removes all items from the cache
func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]Item)
}
