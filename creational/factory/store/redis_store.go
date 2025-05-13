package store

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func ConnectToRedis(addr, port, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password,
		DB:       db,
	})
}

func SaveMapping(short, long string) {
	rdb.Set(ctx, short, long, 0)
}

func RetrieveLongUrl(short string) string {
	res, _ := rdb.Get(ctx, short).Result()
	return res
}