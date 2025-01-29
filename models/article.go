package models

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description int    `json:"description"`
	CateId      string `json:"cate_id"`
	State       int    `json:"state"`
}

func (Article) TableName() string {
	return "article"
}
