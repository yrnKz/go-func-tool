package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()

	mid1 := func(c *gin.Context) {
		fmt.Println("mid1 start")
		test(c)
		fmt.Println("mid1 end")
	}
	mid2 := func(c *gin.Context) {
		fmt.Println("mid2 start")
		fmt.Println("mid2 end")
	}
	router.Use(mid1, mid2)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("process get request")
		c.JSON(http.StatusOK, "hello")
	})
	router.Run()
}
func test(c *gin.Context) {
	c.Next()
}
