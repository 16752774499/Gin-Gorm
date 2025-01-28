package models

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description int    `json:"description"`
	CateId      string `json:"cate_id"`
	State       int    `json:"state"`
	//foreignkey 指定当前表的外键、references 指定关联表中和外键关联的字段
	ArticleCate ArticleCate `gorm:"foreignKey:CateId;references:Id"`
}

func (Article) TableName() string {
	return "article"
}
