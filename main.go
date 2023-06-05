package main

import (
	"github.com/MogLuiz/students-api/api"
)

// type Student struct {
// 	ID       uuid.UUID `json:"id"`
// 	FullName string    `json:"full_name"`
// 	Age      int       `json:"age"`
// }

// var students = []Student{
// 	{ID: shared.GetID(), FullName: "John", Age: 20},
// 	{ID: shared.GetID(), FullName: "Mary", Age: 22},
// }

func main() {
	s := api.New()
	s.Start()
}

// func deleteStudent(c *gin.Context) {
// 	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid id",
// 		})
// 		c.Done()
// 		return
// 	}

// 	for i, s := range students {
// 		if s.ID == id {
// 			students = append(students[:i], students[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "Student deleted",
// 			})
// 			c.Done()
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{
// 		"message": "Student not found",
// 	})
// 	c.Done()
// }

// func showStudent(c *gin.Context) {
// 	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid id",
// 		})
// 		c.Done()
// 		return
// 	}

// 	for _, s := range students {
// 		if s.ID == id {
// 			c.JSON(http.StatusOK, gin.H{
// 				"data": s,
// 			})
// 			c.Done()
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{
// 		"message": "Student not found",
// 	})
// 	c.Done()
// }

// func getRoutes(c *gin.Engine) *gin.Engine {
// 	groupStudents := c.Group("/students")
// 	groupStudents.GET("/", listStudents)
// 	groupStudents.GET("/:id", showStudent)
// 	groupStudents.POST("/", createStudent)
// 	groupStudents.PUT("/:id", updateStudent)
// 	groupStudents.DELETE("/:id", deleteStudent)

// 	return c
// }
