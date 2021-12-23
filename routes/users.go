package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserRoutes(c *gin.Context) {
	c.String(http.StatusOK, "login ")
}

func submitRoutes(c *gin.Context) {
	c.String(http.StatusOK, "submit")
}
