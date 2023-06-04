package main

import (
	"net/http"

	"github.com/MogLuiz/students-api/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swaggo/gin-swagger/example/basic/api"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

var students = []Student{
	{ID: shared.GetID(), FullName: "John", Age: 20},
	{ID: shared.GetID(), FullName: "Mary", Age: 22},
}

func main() {
	s := api.New()
	s.Start()
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
	student.ID = shared.GetID()
	students = append(students, student)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created",
		"data":    student,
	})
	c.Done()
}

func updateStudent(c *gin.Context) {
	var student Student

	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
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
	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
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
	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
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
	groupStudents := c.Group("/students")
	groupStudents.GET("/", listStudents)
	groupStudents.GET("/:id", showStudent)
	groupStudents.POST("/", createStudent)
	groupStudents.PUT("/:id", updateStudent)
	groupStudents.DELETE("/:id", deleteStudent)

	return c
}
