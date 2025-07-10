package controller

import (
	"Tasti-gin/models/request"
	"Tasti-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (c *UserController) Register(ctx *gin.Context) {
	//通过接口获取用户输入的数据，要求使用json数据进行返回
	userInfo := request.RegisterRequest{}
	err := ctx.ShouldBindJSON(userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}
	//下面应该对数据进行处理,传入service层进行处理
	register, err := service.UserServices.Register(userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":   1,
		"userInfo": register,
	})
	return
}

func (c *UserController) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  1,
		"message": "success",
	})
	return
}
