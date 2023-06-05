package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity"
	"github.com/MogLuiz/students-api/entity/shared"
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

	for i, s := range entity.Students {
		if s.ID == id {
			entity.Students[i] = studentDTO
			entity.Students[i].ID = id
			c.JSON(http.StatusOK, gin.H{
				"message": "Student updated",
				"data":    entity.Students[i],
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
