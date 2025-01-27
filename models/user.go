package models

import (
	"time"

	"gorm.io/gorm"
)

// 模型首字母与模型属性首字符大写
type User struct {
	ID        uint           `gorm:"primaryKey"` // 自增主键，类型为无符号整数
	Name      string         `gorm:"size:255"`   // 字符串类型，最大长度 255
	Age       uint8          // 无符号 8 位整数，适合存储较小的年龄值
	Email     string         `gorm:"unique"`   // 字符串类型，且设置为唯一
	Password  string         `gorm:"not null"` // 字符串类型，不为空
	IsActive  bool           // 布尔类型
	BirthDate time.Time      `gorm:"type:date"` // 数据库日期类型
	CreatedAt gorm.DeletedAt `gorm:"index"`     // 带索引的软删除时间字段，记录创建时间（软删除时会有时间戳）
	UpdatedAt gorm.DeletedAt `gorm:"index"`     // 带索引的软删除时间字段，记录更新时间（软删除时会有时间戳）
	DeletedAt gorm.DeletedAt `gorm:"index"`     // 带索引的软删除时间字段，用于软删除标记
}

// 可以使用结构体中自定义方法改变表的名称
func (User) TableName() string {
	return "USER"
}
