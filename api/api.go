package api

import (
	"fmt"

	student_controller "github.com/MogLuiz/students-api/api/controller/students"
	"github.com/MogLuiz/students-api/infra/config"
	"github.com/MogLuiz/students-api/infra/database"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine   *gin.Engine
	Database *database.Database
}

func New(db *database.Database) *Service {
	return &Service{Engine: gin.Default(), Database: db}
}

func (s *Service) GetRoutes() {
	students := s.Engine.Group("/students")
	students.GET("/", student_controller.List)
	students.GET("/:id", student_controller.Show)
	students.POST("/", student_controller.Create)
	students.PUT("/:id", student_controller.Update)
	students.DELETE("/:id", student_controller.Delete)
}

func (s *Service) Start() error {
	s.GetRoutes()
	err := s.Engine.Run(fmt.Sprintf(":%d", config.Env.PORT))
	if err != nil {
		return err
	}

	return nil
}
