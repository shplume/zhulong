package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/zhulong/controller"
)

func Init(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	test := v1.Group("/test")
	test.GET("/ping", controller.Pong)

	user := v1.Group("/user")
	user.POST("/create", controller.CreateUser)
}
