package cache

import (
	"time"

	Cache "github.com/patrickmn/go-cache"
)

var appCache = Cache.New(30*time.Second, 30*time.Second)

// GetCache returns a pointer to the cache
func GetCache() *Cache.Cache {
	return appCache
}
