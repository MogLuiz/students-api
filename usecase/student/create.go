package student

import (
	"github.com/MogLuiz/students-api/entity"
)

func Create(fullName string, age int) (createdStudent *entity.Student) {
	student := entity.CreateStudent(fullName, age)
	entity.Students = append(entity.Students, *student)
	return student
}
