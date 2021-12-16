package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 上传文件
	router := gin.Default()
	// 设置最低限制内存为8MiB
	router.MaxMultipartMemory = 8 << 20

	router.POST("/ipload", func(context *gin.Context) {
		file, _ := context.FormFile("Filename")
		log.Println(file.Filename)

		// 保存上传文件
	})
}
