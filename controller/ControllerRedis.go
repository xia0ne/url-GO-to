package controller

import (
	"ginLearnDemo/model"
	"ginLearnDemo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddUrls struct {
	URL  string  `json:"url"`
	Slug *string `json:"slug"`
}

func AddUrlToRedis(c *gin.Context) {
	req := &AddUrls{}
	err := c.ShouldBindJSON(req)

	if err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Unpredictable error, please contact the administrator",
		})
		return
	}
	var slug string
	if req.Slug != nil {
		slug = *req.Slug
		if exists, err := service.IsExistsInRedis(slug); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Slug already exists, please change your slug",
			})
			return
		} else {
			if exists {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Slug already exists, please change your slug",
				})
				return
			}
		}
		if err := service.SaveUrlToRedis(slug, req.URL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Slug already exists, please change your slug",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": model.MyConfigs.Config.BaseUrl + slug,
		})
	} else {
		if key, err := service.WantPutUrlToRedis(req.URL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Unpredictable error, please contact the administrator",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": model.MyConfigs.Config.BaseUrl + key,
			})
		}
	}
}

func RedirectToUrl(c *gin.Context) {
	key := c.Param("key")
	if url, err := service.GetUrlFromRedis(key); err != nil {
		c.Redirect(http.StatusMovedPermanently, "/err")
	} else {
		c.Redirect(http.StatusMovedPermanently, url)
	}
}
