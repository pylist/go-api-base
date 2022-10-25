package router

import (
	"github.com/gin-gonic/gin"
	"go-api-base/pkg/response"
	"net/http"
)

func Router(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("api")
	api.POST("test", func(c *gin.Context) {
		response.Success(c)
	})
}
