package router

import (
	"Tasti-gin/controller"
	"github.com/gin-gonic/gin"
)

var (
	UserController = controller.UserController{}
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/Register", UserController.Register)
		user.GET("/helloworld", UserController.Test)
	}
	router.Run(":5000")
}
