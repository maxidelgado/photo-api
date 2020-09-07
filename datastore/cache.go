package datastore

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var ds Cache

func init() {
	ds = &localcache{client: cache.New(10*time.Second, 10*time.Second)}
}

func New() Cache {
	return ds
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
	Delete(key string)
}

type localcache struct {
	client *cache.Cache
}

func (l localcache) Get(key string) (interface{}, bool) {
	return l.client.Get(key)
}

func (l localcache) Set(key string, value interface{}, ttl time.Duration) {
	l.client.Set(key, value, ttl)
}

func (l localcache) Delete(key string) {
	l.client.Delete(key)
}
