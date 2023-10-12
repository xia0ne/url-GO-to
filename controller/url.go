package controller

import (
	"ginLearnDemo/model"
	"ginLearnDemo/service"
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

func Add(c *gin.Context, urlstore *service.URLStore, configs *model.Config) {
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
		if !urlstore.Set(slug, req.URL) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Slug already exists",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": configs.Config.BaseUrl + slug,
		})
	} else {
		key := urlstore.Put(req.URL)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": configs.Config.BaseUrl + key,
		})
	}
}

func Myredirect(c *gin.Context, urlstore *service.URLStore) {
	url := c.Param("url")
	key := urlstore.Get(url)
	if key == "" {
		c.Redirect(http.StatusMovedPermanently, "/err")
	}
	c.Redirect(http.StatusMovedPermanently, key)
}
