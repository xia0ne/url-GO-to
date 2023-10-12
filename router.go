package main

import (
	"ginLearnDemo/controller"
	"ginLearnDemo/model"
	"ginLearnDemo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyInitRouter(r *gin.Engine, store *service.URLStore, configs *model.Config) {

	r.Static("/static", "./template/assets")
	r.LoadHTMLGlob("template/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusOK, "404.html", nil)
	})
	r.POST("/add", func(c *gin.Context) {
		controller.Add(c, store, configs)
	})
	router := r.Group("/api")

	router.GET("/:url", func(c *gin.Context) {
		controller.Myredirect(c, store)
	})
}
