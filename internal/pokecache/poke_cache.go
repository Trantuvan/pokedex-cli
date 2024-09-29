package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mu       *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		entries:  map[string]cacheEntry{},
		mu:       &sync.RWMutex{}, //* use pointer ?
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result, ok := c.entries[key]
	return result.val, ok
}

func (c *Cache) reapLoop() {
	timer := time.NewTicker(c.interval)
	defer timer.Stop()

	for range timer.C {
		for k, v := range c.entries {
			c.mu.Lock()
			if delta := time.Since(v.createdAt); delta > c.interval {
				delete(c.entries, k)
			}
			c.mu.Unlock()
		}
	}
}
