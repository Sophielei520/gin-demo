// 浏览器或控制台执行如下命令
// http://localhost:8080/someGet?id=1&page=1

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/someGet", getting)
	router.Run()
}

func getting(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.String(http.StatusOK, "id:%s page:% name:%s message:%s", id, page, name, message)
}
