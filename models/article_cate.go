package models

type ArticleCate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	State int    `json:"state"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
