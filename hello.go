package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/hello", helloWorld)

	router.Run(":8082")
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
