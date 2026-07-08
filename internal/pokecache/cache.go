package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:       sync.Mutex{},
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}

		c.mu.Unlock()
	}
}
