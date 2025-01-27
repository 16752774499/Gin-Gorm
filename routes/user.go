package routes

import (
	"errors"
	"fmt"
	"gin-orm/controller"
	"gin-orm/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestfulUser(ctx *gin.Engine) {
	restfulUser := ctx.Group("/restfuluser")
	restfulUser.GET("/user", UserGet)           // 查看所有，返回所有
	restfulUser.GET("/user/:id", UserGet)       //查看某个具体资源，返回具体
	restfulUser.POST("/user", UserPost)         //添加资源，返回添加资源
	restfulUser.PUT("/user/:id", UserPut)       //编辑某个资源，返回该资源
	restfulUser.DELETE("/user/:id", UserDelete) //删除某个资源，返回空

}
func UserDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	user := models.User{}

	err := models.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
			})
			return
		}

		fmt.Printf("查询记录时出错: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"result":  err,
		})
		return
	}
	// 对数据进行软删除
	result := models.DB.Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		fmt.Printf("软删除失败: %v\n", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"result":  result.Error,
		})
	} else {
		fmt.Printf("软删除成功，影响行数: %d\n", result.RowsAffected)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  fmt.Sprintf("软删除成功，受影响行数:%d", result.RowsAffected),
		})
	}

}

func UserGet(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		//获取全部
		var user []models.User
		models.DB.Find(&user)

		//将结果放进users中
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  user,
		})
	} else {
		//根据条件获取某个

		var user []models.User
		models.DB.Where("id=?", id).Find(&user)
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
	user := models.User{
		Name:      ctx.PostForm("name"),
		Age:       uint8(controller.StringToUint(ctx.PostForm("age"))),
		Email:     ctx.PostForm("email"),
		Password:  ctx.PostForm("password"),
		IsActive:  ctx.PostForm("isActive") == "true",
		BirthDate: time.Now(),
	}

	result := models.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  result.Error,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  user,
		})
	}
}

func UserPut(ctx *gin.Context) {

	id := ctx.Param("id")
	var user models.User
	//将库中数据取出存进user中
	err := models.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
			})
			return
		}
		fmt.Printf("查询记录时出错: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"result":  err,
		})
		return
	}

	var updatedUser models.User
	//将json参数反序列化
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	fmt.Println(updatedUser)
	//更新user
	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Age != 0 {
		user.Age = updatedUser.Age
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}
	user.IsActive = updatedUser.IsActive
	/*
		updatedUser.BirthDate 和 user.BirthDate 假设都是 time.Time 类型，用于表示日期和时间 。
		IsZero() 是 time.Time 类型的一个方法，
		当 time.Time 的值表示零时间（也就是0001-01-01 00:00:00 +0000 UTC ）时，
		IsZero() 方法返回 true，否则返回 false。
	*/
	if !updatedUser.BirthDate.IsZero() {
		user.BirthDate = updatedUser.BirthDate
	}
	//更新库
	if err := models.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error updating user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  user,
	})
}

//SELECT * FROM `USER` WHERE `USER`.`created_at` IS NULL
//SELECT * FROM `USER` WHERE created_at IS NOT NULL AND `USER`.`created_at` IS NULL
