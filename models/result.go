package models

type Result struct {
	ID   int
	Name string
	Age  int
}

func (Result) TableName() string {
	return "RESULT"
}
