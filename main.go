package main

import (
	"gin-orm/models"
	"gin-orm/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var articleList []models.Article
		//Preload("ArticleCate")里面的ArticleCate为Articlestruct中定义的属性ArticleCate
		models.DB.Preload("ArticleCate").Limit(10).Find(&articleList)
		c.JSON(200, gin.H{
			"result": articleList,
		})
	})
	routes.RestfulUser(r)

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}

}
