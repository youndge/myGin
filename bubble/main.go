package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

func main()  {
	

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