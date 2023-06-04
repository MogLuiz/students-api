package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service struct {
	*gin.Engine
}

func New() *Service {
	return &Service{gin.Default()}
}

func (s *Service) GetRoutes() {
	students := s.Engine.Group("/students")
	students.GET("/", listStudents)
}

func (s *Service) Start() error {
	s.GetRoutes()
	err := s.Run(":8080")
	if err != nil {
		return err
	}
	fmt.Println("Starting server on port 8080")
	return nil
}
