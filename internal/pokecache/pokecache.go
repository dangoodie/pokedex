package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]*cacheEntry
	mutex   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]*cacheEntry),
		mutex:   sync.Mutex{},
	}

	// Start the reaploop
	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key *string, val []byte) {
	if key == nil {
		return
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[*key] = &cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key *string) ([]byte, bool) {
	if key == nil {
		return nil, false
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, found := c.entries[*key]
	if found {
		entry.createdAt = time.Now() // reset for a cache hit
		return entry.val, true
	} else {
		return nil, false
	}
}

// Periodically triggers a cache reap
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // will reap every 5 seconds
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

// Reap removes the expired cache entries
func (c *Cache) reap(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > interval {
			delete(c.entries, key)
		}
	}
}
