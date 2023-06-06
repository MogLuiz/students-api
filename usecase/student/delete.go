package student

import (
	"errors"

	"github.com/MogLuiz/students-api/entity"
	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (err error) {
	for i, s := range entity.Students {
		if s.ID == id {
			entity.Students = append(entity.Students[:i], entity.Students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}
