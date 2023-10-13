package test

import (
	"fmt"
	"ginLearnDemo/model"
	"ginLearnDemo/service"
	"testing"
)

func TestSaveRedis(t *testing.T) {
	service.SaveUrlToRedis("123", "123")
}

func TestConfig(t *testing.T) {
	fmt.Println(model.MyConfigs.Redis)
}
