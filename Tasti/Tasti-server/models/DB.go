package models

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/Tasti?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 测试连接
	sqlDB, err := DB.DB()
	if err != nil {
	}
	defer sqlDB.Close()
	// 迁移 schema
	DB.AutoMigrate()

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("数据库Ping失败: %v", err)
	}

	log.Printf("数据库连接成功")
}

// InitRedis 初始化Redis连接
func InitRedis() error {
	// 创建Redis客户端
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码
		DB:       0,  // 默认数据库
		// 连接池配置（建议生产环境添加）
		PoolSize:     100, // 连接池最大连接数
		MinIdleConns: 10,  // 最小空闲连接数
		PoolTimeout:  30,  // 连接池超时时间（秒）
	})

	// 测试连接
	ctx := context.Background()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis连接失败: %v", err)
		return err
	}

	log.Println("Redis连接成功")
	return nil
}
