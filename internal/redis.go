package internal

import (
	"context"
	"crypto-data-collect/global"
	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client {
	redisCfg := global.SERVER_CONFIG.Redis
	client := redis.NewClient(&redis.Options{DB: 0, Password: redisCfg.Password, Addr: redisCfg.Host})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	} else {
		return client
	}
}
