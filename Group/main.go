package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
	}

}

func readEndpoint(c *gin.Context) {
	fmt.Println("read")
}

func loginEndpoint(c *gin.Context) {
	fmt.Println("login")
}
