package cache

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cacheEnv := os.Getenv("CACHE_ADAPTER")
	if cacheEnv == "redis" {
		client = GetRedisCache()
	} else {
		panic("Invalid cache adapter supplied.")
	}
}

func Get(key string) (string, error) {
	value, err := client.Get(key)
	return value, err
}

func Set(key string, value string) (bool, error) {
	result, err := client.Set(key, value)
	return result, err
}
