package service

import (
	"Tasti-gin/dao"
	"Tasti-gin/models"
	"Tasti-gin/models/request"
)

type UserServiceIF interface {
	Register(Userinfo request.RegisterRequest) (models.User, error)
}

type UserService struct {
}

func (s *UserService) Register(Userinfo request.RegisterRequest) (models.User, error) {
	//看是否要对用户信息进行处理，如果不处理就直接传入dao层存入数据库，这里明显不需要传入任何数据
	return dao.UserDo.Register(Userinfo)
}

var UserServices UserServiceIF = &UserService{}
