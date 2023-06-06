package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity/shared"
	student_usecase "github.com/MogLuiz/students-api/usecase/student"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		c.Done()
		return
	}

	s, hasStudent := student_usecase.SearchById(id)
	if !hasStudent {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		c.Done()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
	c.Done()
}
