package students

import (
	"net/http"

	student_usecase "github.com/MogLuiz/students-api/usecase/student"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var dto StudentInputDTO
	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		c.Done()
		return
	}

	student, err := student_usecase.Create(dto.FullName, dto.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Done()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created",
		"data":    student,
	})
	c.Done()
}
