package student

import (
	"github.com/MogLuiz/students-api/entity"
	"github.com/google/uuid"
)

func SearchById(id uuid.UUID) (student entity.Student, hasStudent bool) {
	for _, s := range entity.Students {
		if s.ID == id {
			return s, true
		}
	}
	return student, false
}
