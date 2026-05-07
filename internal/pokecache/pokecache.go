package pokecache

import (
	"time"
	"sync"
	)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) (Cache) {
	c := Cache {
		cache: map[string]cacheEntry{},
	}
	ticker := time.NewTicker(interval)
    go func() {
		for range ticker.C {
		    c.reapLoop(interval)
		}
    }()
	return c
}

func Add(c Cache, key string, val []byte) {
	c.add(key, val)
}

func (c Cache) add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = cacheEntry   { createdAt: time.Now(),
		                          val: val,
                            	}
	c.mu.Unlock()
}

func Get(c Cache, key string) ([]byte, bool) {
	return c.get(key)
}

func (c Cache) get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.cache[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	for key, entry := range c.cache {
		if time.Now().Sub(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
	c.mu.Unlock()
}

