package controller

import (
	"fmt"
	"gin-orm/models"
)

func Add(moudel interface{}) bool {
	result := models.DB.Create(moudel)
	if result.RowsAffected != 0 {
		fmt.Println(moudel, "Add sussecs!")
		return true
	} else {
		fmt.Println(models.DB.Error)
		return false
	}
}

func Delete(moudel interface{}) bool {
	result := models.DB.Delete(moudel)
	if result.RowsAffected != 0 {
		fmt.Println(moudel, "Delete sussecs!")
		return true
	} else {
		fmt.Println(models.DB.Error)
		return false
	}
}

func Updata(moudel interface{}) bool {
	result := models.DB.Updates(moudel)
	if result.RowsAffected != 0 {
		fmt.Println(moudel, "Updata sussecs!")
		return true
	} else {
		fmt.Println(models.DB.Error)
		return false
	}
}
