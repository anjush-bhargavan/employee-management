package config

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisService represents the client of redis.
type RedisService struct {
	Client *redis.Client
}

// SetupRedis initialisez redis server with configuration variables.
func SetupRedis(cfg *Config) (*RedisService, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.REDISHOST,
		DB:   0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New("failed to connect to redis")
	}

	return &RedisService{
		Client: client,
	}, nil
}

// SetDataInRedis will set  data in redis with a key and expiration time.
func (r *RedisService) SetDataInRedis(key string, value []byte, expTime time.Duration) error {
	err := r.Client.Set(context.Background(), key, value, expTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetFromRedis will help to retrieve the data from redis.
func (r *RedisService) GetFromRedis(key string) (string, error) {
	jsonData, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return jsonData, nil
}

// DelFromRedis will help to delete the specified keys from Redis.
func (r *RedisService) DelFromRedis(keys ...string) (int64, error) {
	deleted, err := r.Client.Del(context.Background(), keys...).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
