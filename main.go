package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}

var students = []Student{
	{ID: 1, FullName: "John", Age: 20},
	{ID: 2, FullName: "Mary", Age: 22},
}

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
	c.Done()
}

func listStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List students",
		"data":    students,
	})
	c.Done()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHeart)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", listStudents)

	return c
}
