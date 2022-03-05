package routers

import (
	"win/lab8/controllers"

	"github.com/gin-gonic/gin"
)

func RouterIndex(r *gin.Engine)  {
	r.GET("/",controllers.Index)
}

