package cache_test

import (
	"testing"

	"github.com/dant89/go-gopher-api/src/api/cache"
)

func TestRedisStore(t *testing.T) {
	testString := "value123"

	client := cache.GetRedisCache()
	client.Set("test", testString)

	value, err := client.Get("test")
	if err != nil {
		t.Error("Failed to retrieve stored value.")
		return
	}

	if value != testString {
		t.Error("Failed to validate stored value is correct.")
		return
	}
}
