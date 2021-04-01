package main

import (
	"fmt"
	setting2 "goGinExample/pkg/setting"
	"goGinExample/routers"
	"net/http"
)

func main()  {
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message":"test"})
	//})

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting2.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting2.ReadTimeout,
		WriteTimeout:   setting2.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
