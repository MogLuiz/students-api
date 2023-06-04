package students

import (
	"net/http"

	"github.com/MogLuiz/students-api/entity"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Students)
	c.Done()
}
