package main

import (
	"ginLearnDemo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyInitRouter(r *gin.Engine) {
	r.Static("/static", "./template/assets")
	r.LoadHTMLGlob("template/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusOK, "404.html", nil)
	})

	r.POST("/add", controller.AddUrlToRedis)

	router := r.Group("/api")
	{
		router.GET("/:key", controller.RedirectToUrl)
	}
}
