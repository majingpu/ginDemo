package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "admin",
		})
	})
	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// user/:name/:action return  "john is send"
	// user/:name/*action return  "john is /send"z

	r.GET("user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, "%s is %s", name, action)
	})

	// 自定义http server配置
	/*	s := &http.Server{
			Addr:           ":8080",
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()*/

	r.Run()
}
