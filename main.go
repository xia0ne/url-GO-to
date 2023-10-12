package main

import (
	"ginLearnDemo/model"
	"ginLearnDemo/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	Configs, _ := model.NewConfig().ReadConfig()
	log.Println("the new config: ", Configs.Config.BaseUrl)
	r := gin.Default()
	store := service.NewUrlStore()
	MyInitRouter(r, store, Configs)
	r.Run(Configs.Config.Port)
}
