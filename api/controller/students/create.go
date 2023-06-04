package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity"
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

	student := entity.CreateStudent(dto.FullName, dto.Age)
	entity.Students = append(entity.Students, student)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created",
		"data":    student,
	})
	c.Done()
}
