package main

import (
	"gin-orm/models"
	"gin-orm/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	r := gin.Default()
	//查询学生以及课程信息
	r.GET("/", func(c *gin.Context) {
		studentList := []models.Student{}
		lessonList := []models.Lesson{}
		models.DB.Find(&studentList)
		models.DB.Find(&lessonList)
		c.JSON(200, gin.H{
			"studentList": studentList,
			"lessonList":  lessonList,
		})
	})
	//查询学生以及各自选课结果
	r.GET("/a", func(c *gin.Context) {
		studentList := []models.Student{}
		models.DB.Preload("Lesson").Find(&studentList)
		c.JSON(200, gin.H{
			"studentList": studentList,
		})
	})
	//查询某个学生选修了哪些课程
	r.GET("/find/:id", func(c *gin.Context) {
		student := []models.Student{}
		models.DB.Preload("Lesson").Where("id = ?", c.Param("id")).First(&student)

		c.JSON(200, gin.H{
			"student": student,
		})
	})
	//课程被哪些学生选修了
	r.GET("/c", func(c *gin.Context) {
		lessonList := []models.Lesson{}
		models.DB.Preload("Student").Find(&lessonList)
		c.JSON(200, gin.H{
			"lessonList": lessonList,
		})

	})
	//查询某个课程都被那些学生选修
	r.GET("/d", func(c *gin.Context) {
		lessonList := []models.Lesson{}
		models.DB.Preload("Student").Where("id = ?", 1).Find(&lessonList)
		c.JSON(200, gin.H{
			"lessonList": lessonList,
		})
	})
	//查询指定数据
	r.GET("/e", func(c *gin.Context) {
		lessonList := []models.Lesson{}
		models.DB.Preload("Student").Offset(1).Limit(2).Find(&lessonList)
		c.JSON(200, gin.H{
			"lessonList": lessonList,
		})
	})
	//关联查询指定子集的筛选条件,查询子集时去除相应信息
	r.GET("/f", func(c *gin.Context) {
		lessonList := []models.Lesson{}
		//去除ID为1的数据
		models.DB.Preload("Student", "id!=1").Find(&lessonList)
		c.JSON(http.StatusOK, lessonList)
	})
	//课程被哪些学生选修 要求：学生id倒叙输出
	r.GET("/g", func(c *gin.Context) {
		lessonList := []models.Lesson{}
		//使用作用域也能实现对应功能Scopes()
		//models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		//	return models.DB.Order("id DESC")
		//使用作用域
		models.DB.Scopes(models.CustomPreloadScope).Find(&lessonList)

		c.JSON(http.StatusOK, lessonList)
	})
	routes.RestfulUser(r)

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}

}

// 生成测试数据函数
// 多对测试数据
func GenerateTestData(db *gorm.DB) {
	// 清空旧数据（生产环境慎用）
	db.Exec("DELETE FROM lesson_student")
	db.Exec("DELETE FROM students")
	db.Exec("DELETE FROM lessons")

	// 创建课程
	lessons := []*models.Lesson{
		{Id: 1, Name: "高等数学"},
		{Id: 2, Name: "大学英语"},
		{Id: 3, Name: "计算机基础"},
		{Id: 4, Name: "数据结构"},
		{Id: 5, Name: "算法设计"},
	}
	for _, lesson := range lessons {
		db.Create(lesson)
	}

	// 创建学生
	students := []*models.Student{
		{
			Id:       1,
			Number:   "20230001",
			Password: "mypassword",
			ClassId:  101,
			Name:     "张三",
		},
		{
			Id:       2,
			Number:   "20230002",
			Password: "securepass",
			ClassId:  102,
			Name:     "李四",
		},
		{
			Id:       3,
			Number:   "20230003",
			Password: "test1234",
			ClassId:  101,
			Name:     "王五",
		},
		{
			Id:       4,
			Number:   "20230004",
			Password: "studentpw",
			ClassId:  103,
			Name:     "赵六",
		},
		{
			Id:       5,
			Number:   "20230005",
			Password: "helloworld",
			ClassId:  102,
			Name:     "陈七",
		},
	}
	for _, student := range students {
		db.Create(student)
	}

	// 建立关联关系（使用GORM的Association）
	db.Model(students[0]).Association("Lesson").Append([]*models.Lesson{
		lessons[0], lessons[1], lessons[2],
	})

	db.Model(students[1]).Association("Lesson").Append([]*models.Lesson{
		lessons[0], lessons[3], lessons[4],
	})

	db.Model(students[2]).Association("Lesson").Append([]*models.Lesson{
		lessons[1], lessons[2], lessons[3],
	})

	db.Model(students[3]).Association("Lesson").Append([]*models.Lesson{
		lessons[0], lessons[2], lessons[4],
	})

	db.Model(students[4]).Association("Lesson").Append([]*models.Lesson{
		lessons[1], lessons[3], lessons[4],
	})
}
