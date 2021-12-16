package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// form 表达形式提交数据
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.POST("/test", message1)

	router.Run()

}

func message1(context *gin.Context) {
	message := context.PostForm("message")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
