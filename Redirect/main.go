// 浏览器或控制台执行如下命令
// http://localhost:8080/someGet?ids[a]=1234&ids[b]=hello

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/raw", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://baidu.com")
	})
	router.Run(":8080")
}
