package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	value     []byte
	createdAt time.Time
}

func NewCache(cacheTimeToLive time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.cacheReaper(cacheTimeToLive)
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.value, ok
}

func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{
		value:     value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) cacheReaper(timeToLive time.Duration) {
	ticker := time.NewTicker(timeToLive)
	for range ticker.C {
		c.reap(timeToLive)
	}
}

func (c *Cache) reap(timeToLive time.Duration) {
	for key, entry := range c.cache {
		if entry.createdAt.Before(time.Now().UTC().Add(-timeToLive)) {
			delete(c.cache, key)
		}
	}
}
