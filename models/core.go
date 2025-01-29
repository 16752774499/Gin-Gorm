package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB
var err error

// 初始化函数默认执行没有参数没有返回值
func init() {
	dsn := "gorm:ECweAtSJPaSBffd3@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             200,         // 慢查询阈值
				LogLevel:                  logger.Info, // 日志级别，设置为Info将打印详细信息
				IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
				Colorful:                  true,        // 是否彩色打印
			},
		),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("初始化！")
	//自动迁移表结构
	err = DB.AutoMigrate(&User{}, &ArticleCate{}, &Article{}, &LessonStudent{}, &Lesson{}, &Student{})
	if err != nil {
		fmt.Printf("Failed to auto migrate: %v\n", err)
	}

}
