package service

import (
	"ginLearnDemo/model"
	"ginLearnDemo/utils"
	"log"
	"time"
)

func SaveUrlToRedis(key, value string) error {
	times, _ := time.ParseDuration(model.MyConfigs.Redis.CacheDuration)
	err := model.MyRedis.RedisClient.Set(model.Ctx, key, value, times).Err()
	if err != nil {
		return err
	}
	countKey := key + ":count"
	err = model.MyRedis.RedisClient.Incr(model.Ctx, countKey).Err()
	if err != nil {
		return err
	}
	return nil
}

func IsExistsInRedis(key string) (bool, error) {
	if exists, err := model.MyRedis.RedisClient.Exists(model.Ctx, key).Result(); err != nil {
		return false, err
	} else {
		return exists == 1, nil
	}
}

func WantPutUrlToRedis(key string) (string, error) {
	slug := utils.Base62()
	if err := SaveUrlToRedis(slug, key); err != nil {
		return "", err
	} else {
		return slug, nil
	}
}

func IncrKeysToRedis(key string) error {
	keyss := key + ":count"
	if err := model.MyRedis.RedisClient.Incr(model.Ctx, keyss).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUrlFromRedis(key string) (string, error) {
	result, err := model.MyRedis.RedisClient.Get(model.Ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func GetUrlCountsFromRedis(key string) (int64, error) {
	countKey := key + ":count"
	if ans, err := model.MyRedis.RedisClient.Get(model.Ctx, countKey).Int64(); err != nil {
		return -1, err
	} else {
		return ans, nil
	}

}

func GetUrlTTLFromRedis(key string) (time.Duration, error) {
	if ans, err := model.MyRedis.RedisClient.TTL(model.Ctx, key).Result(); err != nil {
		return -2, nil
	} else {
		return ans, nil
	}
}
