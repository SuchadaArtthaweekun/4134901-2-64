package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context)  {
	id := c.Param("id")

	c.JSON(
		http.StatusOK,gin.H{
			"id":id,
		},
	)
}

func DelAllUser(c *gin.Context)  {
	c.JSON(
		http.StatusOK,gin.H{
			"title":"Delete All Users",
		},
	)
}