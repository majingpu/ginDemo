package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addPingRoutes(r *gin.RouterGroup) {
	ping := r.Group("/ping")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
