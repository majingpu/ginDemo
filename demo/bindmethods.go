package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	/*
		form: 请求体中包含表单
		json: 请求体中包含json
		xml: 请求体中包含xml
		binding: required 指定User为必须绑定的字段
	*/
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {

	// binding 将请求体绑定到结构体
	router := gin.Default()

	// json 使用ShouldBindJSON
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/login/json", func(context *gin.Context) {
		var json Login

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if json.User != "admin" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})

	//  Example for binding XML 使用 ShouldBindXML
	/*
		<?xml version="1.0" encoding="UTF-8"?>
		<root>
		<user>admin</user>
		<password>123</password>
		</root>
	*/
	router.POST("login/xml", func(context *gin.Context) {
		var xml Login
		if err := context.ShouldBindXML(&xml); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if xml.User != "admin" || xml.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})

	// form 使用 ShouldBind
	router.POST("login/form", func(context *gin.Context) {
		var form Login
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if form.User != "admin" || form.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})

	// ShouldBindQuery 如果 url 查询参数和 post 数据都存在，函数只绑定 url 查询参数而忽略 post 数据
	// login/url?user=admin&password=123
	router.GET("/login/url", func(context *gin.Context) {
		var url Login

		if err := context.ShouldBindQuery(&url); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "hello world",
		})

	})

	// running
	router.Run()
}
