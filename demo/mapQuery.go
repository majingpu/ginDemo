package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	/*
		parameter: /post?ids[1]=12&ids[2]=14
		return: ids: map[1:12 2:14];

		form-data:  key:names[first],value:test1
		return: names: map[first:test1]
	*/
	router.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	router.Run()
}
