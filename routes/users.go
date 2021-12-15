package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})

	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})

	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
