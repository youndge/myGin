package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

func initMySQL() (err error){
	//数据库链接信息dsn
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil{
		fmt.Printf("err:%v",err)
		return
	}
	err = DB.DB().Ping()
	return
}

func main()  {
	//创建数据库
	//sql:CREATE DATEBASE bubble;
	//链接数据库
	err := initMySQL()
	if err != nil{
		panic(err)
	}
	defer DB.Close()//程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&Todo{})


	r := gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/",func(c *gin.Context)  {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//代办事项
		//ADD
		v1Group.POST("/todo",func(c *gin.Context)  {
			//前端页面填写待办事项，点击提交，会发送请求到这里
			//1.从请求中把数据取出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			//DB.Create(&todo)
			//3.返回响应
			if err = DB.Create(&todo).Error;err != nil{
				c.JSON(http.StatusOK, gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,todo)
			}
		})
		//查看所有代办事项
		v1Group.GET("/todo",func(c *gin.Context)  {
			
			
		})
		//查看某一个代办事项
		v1Group.GET("/todo:id",func(c *gin.Context)  {
			
		})
		//查看所有代办事项
		v1Group.PUT("/todo:id",func(c *gin.Context)  {
			
		})
		//查看所有代办事项
		v1Group.DELETE("/todo:id",func(c *gin.Context)  {
			
		})

	}

	r.Run()
}