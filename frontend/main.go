package main

import (
	"win/lab8/configs"
	"win/lab8/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Connection()

	r := gin.Default()

	routers.RouterIndex(r)
	routers.RouterUser(r)

	r.Run(":8000")
}