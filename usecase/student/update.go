package student

import (
	"errors"

	"github.com/MogLuiz/students-api/entity"
	"github.com/google/uuid"
)

func Update(id uuid.UUID, studentDTO entity.Student) (err error, updatedStudent entity.Student) {
	for i, s := range entity.Students {
		if s.ID == id {
			entity.Students[i] = studentDTO
			entity.Students[i].ID = id
			return nil, entity.Students[i]
		}
	}
	return errors.New("student not found"), entity.Student{}
}
