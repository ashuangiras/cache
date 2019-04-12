package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

var CacheStore *cache.Cache

func InitCacheStore() {
	CacheStore = cache.New(5*time.Minute, 10*time.Minute)
	log.Println("Successfully initialized the Cache Store")
}

func StoreInCache(key string, val []byte) {
	CacheStore.Set(key, val, cache.DefaultExpiration)
}

func LookInCache(key string) (string, bool) {
	data, ok := CacheStore.Get(key)
	if !ok {
		return "", false
	}
	return fmt.Sprintf("%s", data), ok
}
