package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func main() {

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "rc-yym-dev-usce1.redis.cache.windows.net:6379",
		Password: os.Getenv("redisKey"),
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "super cool cached message", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
