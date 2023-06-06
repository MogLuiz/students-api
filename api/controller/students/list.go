package students

import (
	"net/http"

	student_usecase "github.com/MogLuiz/students-api/usecase/student"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Done()
		return
	}
	c.JSON(http.StatusOK, students)
	c.Done()
}
