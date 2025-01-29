package models

type ArticleCate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	State int    `json:"state"`
	//一对多数据模型
	Article []Article `gorm:"foreignKey:CateId"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
