package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addPingRoutes(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}
