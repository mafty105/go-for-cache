package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	_interface "go-testing-for-cache/application/interface"
	"time"
)

type RedisCache[T any] struct {
	client    *redis.Client
	keyPrefix string
}

func NewRedisCache[T any](client *redis.Client, keyPrefix string) *RedisCache[T] {
	return &RedisCache[T]{
		client:    client,
		keyPrefix: keyPrefix,
	}
}

func (rc *RedisCache[T]) formatKey(key string) string {
	return fmt.Sprintf("%s%s", rc.keyPrefix, key)
}

func (rc *RedisCache[T]) Get(ctx context.Context, key string) (*T, error) {
	formattedKey := rc.formatKey(key)

	val, err := rc.client.Get(ctx, formattedKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, _interface.ErrCacheMiss
		}
		return nil, fmt.Errorf("キャッシュからの読み取りエラー: %w", err)
	}

	var result T
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, fmt.Errorf("JSONデコードエラー: %w", err)
	}

	return &result, nil
}

func (rc *RedisCache[T]) Set(ctx context.Context, key string, value T, ttl time.Duration) error {
	formattedKey := rc.formatKey(key)

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("JSONエンコードエラー: %w", err)
	}

	err = rc.client.Set(ctx, formattedKey, data, ttl).Err()
	if err != nil {
		return fmt.Errorf("キャッシュへの書き込みエラー: %w", err)
	}

	return nil
}

func (rc *RedisCache[T]) Delete(ctx context.Context, key string) error {
	formattedKey := rc.formatKey(key)

	err := rc.client.Del(ctx, formattedKey).Err()
	if err != nil {
		return fmt.Errorf("キャッシュからの削除エラー: %w", err)
	}

	return nil
}
