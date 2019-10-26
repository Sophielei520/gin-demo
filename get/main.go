// 浏览器或控制台执行如下命令
// http://localhost:8080/user

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "zhangsan",
		})
	})
	router.Run()
}
