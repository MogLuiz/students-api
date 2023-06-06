package student

import (
	"errors"

	"github.com/MogLuiz/students-api/entity"
	"github.com/google/uuid"
)

func Create(fullName string, age int) (createdStudent *entity.Student, err error) {
	student := entity.CreateStudent(fullName, age)
	entity.Students = append(entity.Students, *student)

	if student.ID == uuid.Nil {
		return nil, errors.New("internal server error")
	}

	return student, nil
}
