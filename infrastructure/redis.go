package infrastructure

import (
	"github.com/go-redis/redis/v8"
)

func OpenRedis() *redis.Client  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "p4ssw0rd", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
