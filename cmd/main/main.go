package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shplume/zhulong/logger"
	"github.com/shplume/zhulong/routers"
)

func main() {
	r := gin.Default()

	routers.Init(r)

	r.Run()
}
