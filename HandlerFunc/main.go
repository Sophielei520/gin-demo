package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	//使用日志
	router.Use(gin.Logger())
	//使用Panic处理方案
	router.Use(gin.Recovery())
	//路由途径多个函数
	router.GET("/index", index(), chain)

	router.Run(":8080")
}

func index() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
	}
}
func chain(c *gin.Context) {
	fmt.Println("index chainFun")
}
