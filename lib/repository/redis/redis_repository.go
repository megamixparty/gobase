package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/megamixparty/gobase/lib/repository"
)

// RedisKey store prefix of Redis key for this service
const RedisKey = "gobase"

// RedisRepository store client connection to Redis
type RedisRepository struct {
	redis *redis.Client
}

// NewRedisRepository returns database implementation of IMemoryRepository
func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{redis}
}

// Get return value in redis based on key given
func (rr *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	queryKey := fmt.Sprintf("%s:%s", RedisKey, key)

	result, err := rr.redis.Get(ctx, queryKey).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

// Set assign value to redis based on key given
func (rr *RedisRepository) Set(ctx context.Context, key string, value string) error {
	queryKey := fmt.Sprintf("%s:%s", RedisKey, key)

	err := rr.redis.Set(ctx, queryKey, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// SetEx assign value to redis based on key given
func (rr *RedisRepository) SetEx(ctx context.Context, key string, value string, duration time.Duration) error {
	queryKey := fmt.Sprintf("%s:%s", RedisKey, key)

	err := rr.redis.Set(ctx, queryKey, value, duration).Err()
	if err != nil {
		return err
	}

	return nil
}

// NewTokenRepository return TokenRepository
func (rr *RedisRepository) NewTokenRepository() repository.ITokenRepository {
	return NewTokenRedis(rr.redis)
}
