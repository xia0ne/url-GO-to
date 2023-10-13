package main

import (
	"ginLearnDemo/model"
	"ginLearnDemo/service"
	"github.com/gin-gonic/gin"
)

func init() {
	model.MyConfigs = service.ReadConfig()
	model.MyRedis = service.InitializeStore()
	model.MyStore = model.NewUrlStore()
}

func main() {
	r := gin.Default()
	MyInitRouter(r)
	r.Run(model.MyConfigs.Config.Port)
}
