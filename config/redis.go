package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql" // Mysql driver
)

// NewRedis create new redis database instance and return it
func NewRedis(cfg *EnvConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	return rdb, err
}
