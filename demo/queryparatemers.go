package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	/*
	/welcome?firstname=1   return:hello 1
	/welcome               return:hello Guest
	/welcome?lastname=2    return:Hello Guest 2
	/welcome?firstname=1&lastname=2    return:hello 1 2

	*/
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname","Guest") // 返回查询参数的值，没有参数则返回默认值
		lastname := context.Query("lastname")
		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run()
}
