package entity

import (
	"github.com/MogLuiz/students-api/entity/shared"
	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

var Students = []Student{
	{ID: shared.GetID(), FullName: "John", Age: 20},
	{ID: shared.GetID(), FullName: "Mary", Age: 22},
}

func CreateStudent(fullName string, age int) *Student {
	return &Student{
		ID:       shared.GetID(),
		FullName: fullName,
		Age:      age,
	}
}
