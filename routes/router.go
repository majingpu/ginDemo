package routes

import "github.com/gin-gonic/gin"

var r = gin.Default()

func Run() {
	r.Run()

}

func getRoutes() {
	v1 := r.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)
	v2 := r.Group("/v2")
	addPingRoutes(v2)
}
