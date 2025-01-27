package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// 初始化函数默认执行没有参数没有返回值
func init() {
	dsn := "gorm:ECweAtSJPaSBffd3@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("初始化！")
	// 自动迁移表结构
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("Failed to auto migrate: %v\n", err)
	}

}
