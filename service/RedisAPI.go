package service

import (
	"ginLearnDemo/model"
	"ginLearnDemo/utils"
	time2 "time"
)

func SaveUrlToRedis(key, value string) error {
	times, _ := time2.ParseDuration(model.MyConfigs.Redis.CacheDuration)
	err := model.MyRedis.RedisClient.Set(model.Ctx, key, value, times).Err()
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

func GetUrlFromRedis(key string) (string, error) {
	result, err := model.MyRedis.RedisClient.Get(model.Ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
