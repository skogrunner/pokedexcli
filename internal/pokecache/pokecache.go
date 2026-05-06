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

func dummy() int {
	return 0
}