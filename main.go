package main

import (
	"fmt"
	"net/http"
	"strconv"

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
		"data": students,
	})
	c.Done()
}

func createStudent(c *gin.Context) {
	var student Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		c.Done()
		return
	}
	student.ID = students[len(students)-1].ID + 1
	students = append(students, student)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created",
		"data":    student,
	})
	c.Done()
}

func updateStudent(c *gin.Context) {
	var student Student

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		c.Done()
		return
	}

	err = c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		c.Done()
		return
	}

	for i, s := range students {
		if s.ID == id {
			students[i] = student
			students[i].ID = id
			c.JSON(http.StatusOK, gin.H{
				"message": "Student updated",
				"data":    students[i],
			})
			c.Done()
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Student not found",
	})
	c.Done()
}

func deleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		c.Done()
		return
	}

	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Student deleted",
			})
			c.Done()
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Student not found",
	})
	c.Done()
}

func showStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		c.Done()
		return
	}

	for _, s := range students {
		if s.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": s,
			})
			c.Done()
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Student not found",
	})
	c.Done()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHeart)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", listStudents)
	groupStudents.GET("/:id", showStudent)
	groupStudents.POST("/", createStudent)
	groupStudents.PUT("/:id", updateStudent)
	groupStudents.DELETE("/:id", deleteStudent)

	return c
}
