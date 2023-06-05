package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity"
	"github.com/MogLuiz/students-api/entity/shared"
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

	for _, s := range entity.Students {
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
