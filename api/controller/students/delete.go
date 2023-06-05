package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity"
	"github.com/MogLuiz/students-api/entity/shared"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id, err := shared.ConvertStringToUUId(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		c.Done()
		return
	}

	for i, s := range entity.Students {
		if s.ID == id {
			entity.Students = append(entity.Students[:i], entity.Students[i+1:]...)
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
