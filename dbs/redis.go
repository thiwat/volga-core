package dbs

import (
	"context"
	"time"
	"volga-core/configs"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client = ConnectRedis()

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: configs.GetRedisURI(),
		DB:   configs.GetRedisDB(),
	})

	return rdb
}

func SetKey(key string, value string, ttl int) error {
	ctx := context.Background()

	err := RedisClient.Set(ctx, key, value, time.Duration(ttl)*time.Minute).Err()

	return err
}

func GetKey(key string) (string, error) {
	ctx := context.Background()

	val, err := RedisClient.Get(ctx, key).Result()

	return val, err
}
