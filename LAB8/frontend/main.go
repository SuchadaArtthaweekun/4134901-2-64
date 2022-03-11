package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.run(":8000")
}

