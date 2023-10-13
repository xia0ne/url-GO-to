package model

import (
	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	RedisClient *redis.Client
}
