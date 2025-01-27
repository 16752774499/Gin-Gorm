package routes

import (
	"fmt"
	"gin-orm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RestfulUser(ctx *gin.Engine) {
	restfulUser := ctx.Group("/restfuluser")
	restfulUser.GET("/user", UserGet)       // 查看所有，返回所有
	restfulUser.GET("/user/:name", UserGet) //查看某个具体资源，返回具体
	restfulUser.POST("/user", UserPost)     //添加资源，返回添加资源
	restfulUser.PUT("/user", UserPut)       //编辑某个资源，返回该资源
	restfulUser.DELETE("/user", UserDelete) //删除某个资源，返回空

}
func UserDelete(ctx *gin.Context) {

}

func UserGet(ctx *gin.Context) {
	if models.DB == nil {
		fmt.Println("数据库连接未初始化")
	}
	name := ctx.Param("name")
	if name == "" {
		//获取全部
		user := []models.User{}
		fmt.Println(user)
		models.DB.Find(&user)
		//将结果放进users中
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  user,
		})
	} else {
		//根据条件获取某个
		user := []models.User{}
		models.DB.Where("name=?", name).Find(&user)
		if len(user) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  "不存在该条记录！",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  user,
			})
		}

	}

}

func UserPost(ctx *gin.Context) {

}

func UserPut(ctx *gin.Context) {

}
