package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity"
	"github.com/MogLuiz/students-api/entity/shared"
	student_usecase "github.com/MogLuiz/students-api/usecase/student"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var studentDTO entity.Student

	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		c.Done()
		return
	}

	if err := c.BindJSON(&studentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		c.Done()
		return
	}

	err, updatedStudent := student_usecase.Update(id, studentDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		c.Done()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated",
		"data":    updatedStudent,
	})
	c.Done()

}
