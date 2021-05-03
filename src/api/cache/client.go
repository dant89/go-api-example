package cache

type CacheClientOperations interface {
	Get(key string) (string, error)
	Set(key string, value string) (bool, error)
}

type CacheClient struct{}

var client CacheClient
