package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	new_cache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}
	go new_cache.reapLoop(interval)
	return new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}

	return entry.value, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Entries {
			if entry.createdAt.Add(interval).Before(time.Now()) {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}
