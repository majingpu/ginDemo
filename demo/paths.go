package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Disable log's color
	gin.DisableConsoleColor()

	router := gin.Default()

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name") //context.Param
		context.String(http.StatusOK, "return %s", name)
	})

	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		context.String(http.StatusOK, "return name:%s, action:%s", name, action)
	})

	router.POST("/user/:name/*action", func(context *gin.Context) {
		paths := context.FullPath() // return path full path 返回匹配到的完整路径(string)
		if paths == "/user/:name/*action" {
			context.String(http.StatusOK, "%t", paths)
		}
	})

	router.GET("/user/groups", func(context *gin.Context) {
		context.String(http.StatusOK, "The available groups are [...]")
	})

	router.Run(":8080")
}
