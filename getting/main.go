// 浏览器或控制台执行如下命令
// http://localhost:8080/user

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user", getting)
	router.Run()
}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "zhangsan")
}
