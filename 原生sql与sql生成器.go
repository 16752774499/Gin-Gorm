package main

//import (
//	"database/sql"
//	"fmt"
//	"gin-orm/models"
//	"gorm.io/gorm"
//)
//
//func main() {
//	//原生sql查询
//	user := models.User{}
//	user1 := models.User{}
//	models.DB.Raw("SELECT id,name ,age FROM USER  WHERE  id = ?", 34).Scan(&user)
//	fmt.Println(user)
//	models.DB.Raw("SELECT id, name, age FROM USER WHERE name = ?", "齐波").Scan(&user1)
//	fmt.Println(user1)
//	var age int
//	models.DB.Raw("SELECT SUM(age) FROM USER WHERE birth_date= ?", "2025-01-27").Scan(&age)
//	fmt.Println(age)
//	//Exec原生sql
//	models.DB.Exec("UPDATE USER SET age = ? WHERE id IN ?", 1, []int64{29, 30, 31})
//	//命名参数
//	//GORM 支持 sql.NamedArg、map[string]interface{}{} 或 struct 形式的命名参数
//	user3 := models.User{}
//	models.DB.Where("age = @age AND name = @name", sql.Named("age", "22"), sql.Named("name", "齐波")).Find(&user3)
//	fmt.Println(user3)
//	user4 := models.User{}
//	type NamedArgument struct {
//		Age  int
//		Name string
//	}
//	models.DB.Raw("SELECT * FROM USER WHERE (name = @Name AND age = @Age)",
//		NamedArgument{Name: "齐波", Age: 22}).Find(&user4)
//	fmt.Println(user4)
//	//DryRun 模式（干运行）
//	//在不执行的情况下生成 SQL 及其参数，可以用于准备或测试生成的 SQL
//	user5 := models.User{}
//	/*获取一个用于查询的 Statement 对象
//	.Session 方法用于创建一个新地会话，允许在当前数据库连接上进行一些自定义配置操作。
//	这里传入了一个配置对象 &gorm.Session{DryRun: true}，其中 DryRun 设置为 true 表示开启 “干运行” 模式。
//	在干运行模式下，GORM 不会真正执行数据库操作（例如不会实际执行 INSERT、UPDATE、DELETE 等语句），
//	而是生成相应的 SQL 语句（但是实际生成的 SQL 不会执行到数据库中），用于调试和查看即将执行的 SQL 语句是否符合预期。
//	*/
//	stmt := models.DB.Session(&gorm.Session{DryRun: true}).First(&user5, 30).Where("age = ?", 1).Statement
//	fmt.Println("RryRun模式")
//	fmt.Println("生成的 SQL:", stmt.SQL.String())
//	fmt.Println("参数:", stmt.Vars)
//	//ToSQL,返回生成的sql但不执行
//	fmt.Println("ToSQL模式")
//	user6 := models.User{}
//	sqlStr := models.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
//		return tx.Where("age = ?", 1).First(&user6, 30)
//	})
//	fmt.Println("生成的 SQL::", sqlStr)
//}
