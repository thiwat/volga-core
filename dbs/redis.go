package dbs

import (
	"context"
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

func SetKey(key string, value string) error {
	ctx := context.Background()

	err := RedisClient.Set(ctx, key, value, 0).Err()

	return err
}
