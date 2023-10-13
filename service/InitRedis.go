package service

import (
	"fmt"
	"ginLearnDemo/model"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func InitializeStore() *model.StorageService {
	myDB, err := strconv.ParseInt(model.MyConfigs.Redis.DB, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     model.MyConfigs.Redis.Addr,
		Password: model.MyConfigs.Redis.Passwd,
		DB:       int(myDB),
	})

	pong, err := redisClient.Ping(model.Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	return &model.StorageService{
		RedisClient: redisClient,
	}
}
