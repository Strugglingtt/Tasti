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
	//注意gin区分路由大小写
	user := router.Group("/user")
	{
		user.POST("/register", UserController.Register)
		user.GET("/helloworld", UserController.Test)
	}
	router.Run(":5000")
}
