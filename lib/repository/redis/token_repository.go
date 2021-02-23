package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const TokenKey = "token"

// TokenRedis store redis client
type TokenRedis struct {
	redis *redis.Client
}

// NewTokenRedis return TokenRedis
func NewTokenRedis(redis *redis.Client) *TokenRedis {
	return &TokenRedis{redis}
}

// GetToken return token based on userId
func (tk *TokenRedis) GetToken(ctx context.Context, userID int64) (string, error) {
	queryKey := fmt.Sprintf("%s:%s:%d", RedisKey, TokenKey, userID)

	result, err := tk.redis.Get(ctx, queryKey).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

// SetToken return token based on userId
func (tk *TokenRedis) SetToken(ctx context.Context, userID int64, token string, duration time.Duration) error {
	queryKey := fmt.Sprintf("%s:%s:%d", RedisKey, TokenKey, userID)

	err := tk.redis.Set(ctx, queryKey, token, duration).Err()
	if err != nil {
		return err
	}

	return nil
}
