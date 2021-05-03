package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedisCache() CacheClient {
	cache := CacheClient{}
	return cache
}

func (c CacheClient) Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c CacheClient) Set(key string, value string) (bool, error) {
	set, err := rdb.SetNX(ctx, key, value, 10*time.Second).Result()
	if err != nil {
		return false, err
	}
	return set, nil
}
