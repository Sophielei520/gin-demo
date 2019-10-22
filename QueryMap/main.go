// 浏览器或控制台执行如下命令
// http://localhost:8080/someGet?ids[a]=1234&ids[b]=hello

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/someGet", getting)
	router.Run()
}

func getting(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("name")

	fmt.Println("ids:%v; names:%v", ids, names)
	c.JSON(200, gin.H{
		"ids":   ids,
		"names": names,
	})

	c.String(http.StatusOK, "ids:%v;names:%v", ids, names)
}
