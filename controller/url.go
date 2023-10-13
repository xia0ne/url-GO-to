package controller

import (
	"ginLearnDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

func HomePage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, AddForm)
}

func GetCurrentURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil { // 检查是否启用了 TLS/SSL
		scheme = "https"
	}
	host := c.Request.Host
	url := scheme + "://" + host + c.Request.URL.String() + "/"
	return url
}

type AddRequest struct {
	URL  string  `json:"url"`
	Slug *string `json:"slug"`
}

func Add(c *gin.Context) {
	var req AddRequest
	err := c.ShouldBindJSON(&req)
	if err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "unknown",
		})
		return
	}
	var slug string
	if req.Slug != nil {
		slug = *req.Slug
		if !model.MyStore.Set(slug, req.URL) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Slug already exists",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			//"message":
			"message": model.MyConfigs.Config.BaseUrl + slug,
		})

	} else {
		key := model.MyStore.Put(req.URL)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": model.MyConfigs.Config.BaseUrl + key,
		})
	}
}

func Myredirect(c *gin.Context) {
	url := c.Param("keys")
	key := model.MyStore.Get(url)
	if key == "" {
		c.Redirect(http.StatusMovedPermanently, "/err")
	}
	c.Redirect(http.StatusMovedPermanently, key)
}
