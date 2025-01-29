package main

import (
	"gin-orm/models"
	"gin-orm/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		articleCateList := []models.ArticleCate{}
		models.DB.Preload("Article").Find(&articleCateList)
		c.JSON(200, gin.H{
			"result": articleCateList,
		})
	})

	routes.RestfulUser(r)

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}

}
