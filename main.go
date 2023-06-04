package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
	r := gin.Default()
	r = getRoutes(r)
	r.Run(":8080")
}

func routeHeart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHeart)

	return c
}
