package dao

import (
	"Tasti-gin/models"
	"Tasti-gin/models/request"
	"gorm.io/gorm"
)

type UserDaoIF interface {
	Register(Userinfo request.RegisterRequest) (models.User, error)
}

type UserDao struct{}

func (d *UserDao) Register(info request.RegisterRequest) (models.User, error) {
	// 创建用户对象并映射字段
	user := models.User{
		Username: info.Username,
		Password: info.Password, // 注意：这里应该先对密码加密
		// 其他需要映射的字段...
	}

	// 执行数据库插入操作
	result := models.DB.Create(&user)

	// 检查是否有错误发生
	if result.Error != nil {
		return models.User{}, result.Error
	}

	// 检查是否插入成功（影响行数为1）
	if result.RowsAffected == 0 {
		return models.User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}

var UserDo UserDaoIF = &UserDao{}
