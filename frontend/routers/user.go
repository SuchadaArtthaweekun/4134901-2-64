package routers

import (
	"win/lab8/controllers"

	"github.com/gin-gonic/gin"
)

func RouterUser(r *gin.Engine)  {
	
	r.POST("/users/:id",controllers.AddUser)
	r.DELETE("/users",controllers.DelAllUser)
}
