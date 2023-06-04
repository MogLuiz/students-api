package shared

import (
	"log"

	"github.com/google/uuid"
)

func GetID() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Fatal Error", err)
	}
	return uuid
}

func ConvertStringToUUId(id string) (uuid.UUID, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
