package cachehandler

import (
	"sync"
	"time"
	"github.com/patrickmn/go-cache"
)

var (
	C    *cache.Cache
	once sync.Once
)

// InitializeCache ensures that the cache is initialized only once
func InitializeCache() {
	once.Do(func() {
		C = cache.New(2*time.Minute, 10*time.Minute)
	})
}
