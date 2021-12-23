package routes

import "github.com/gin-gonic/gin"

var r = gin.Default()

func Run() {
	r.Run()
}

func getRoutes() {
	v1 := r.Group("/v1")
	{
		v1.GET("/login", addUserRoutes)
		v1.GET("/submit", submitRoutes)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", addPingRoutes)
	}
}
