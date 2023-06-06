package student

import (
	"github.com/MogLuiz/students-api/entity"
)

func List() (students []entity.Student, err error) {
	return entity.Students, nil
}
