package repository

import (
	"context"
	"time"
)

// IMemoryRepository interface
type IMemoryRepository interface {
	// Non-bounded function
	Get(context.Context, string) (string, error)
	Set(context.Context, string, string) error
	SetEx(context.Context, string, string, time.Duration) error

	// Object-based function
	NewTokenRepository() ITokenRepository
}

// ITokenRepository interface
type ITokenRepository interface {
	GetToken(context.Context, int64) (string, error)
	SetToken(context.Context, int64, string, time.Duration) error
}
