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

func (c Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry   { createdAt: time.Now(),
		                          val: val,
                            	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c.cache[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	for key, entry := range c.cache {
		if time.Now().Sub(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}

