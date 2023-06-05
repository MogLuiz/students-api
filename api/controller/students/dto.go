package students

import "github.com/google/uuid"

type StudentInputDTO struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}
