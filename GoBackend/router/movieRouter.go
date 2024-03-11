package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//healthcheck endpoint
	r.GET("/movieapi/v1/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	return r
}
