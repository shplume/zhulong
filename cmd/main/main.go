package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/zhulong/routers"
)

func main() {
	r := gin.Default()

	routers.Init(r)

	r.Run()
}
