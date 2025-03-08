package _interface

import (
	"context"
	"errors"
	"time"
)

var ErrCacheMiss = errors.New("cache: key not found")

// Cache はキャッシュ操作の汎用インターフェースを定義します
type Cache[T any] interface {
	Get(ctx context.Context, key string) (*T, error)
	Set(ctx context.Context, key string, value T, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}
