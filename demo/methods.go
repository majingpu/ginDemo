package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/someGet", getting)
		v1.POST("/somePost", posting)
		v1.PUT("/somePut", putting)
	}

	router.Run(":8080")
}

func putting(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"test": "123",
	})
}

func posting(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"account": "admin",
		"pass":    "12",
	})
}

func getting(context *gin.Context) {
	context.String(http.StatusOK, "Get Demo")
}
