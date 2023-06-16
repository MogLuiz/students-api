package main

import (
	"context"
	"log"

	"github.com/MogLuiz/students-api/api"
	"github.com/MogLuiz/students-api/infra/config"
	"github.com/MogLuiz/students-api/infra/database"
	"github.com/MogLuiz/students-api/infra/database/mongo"
)

func main() {
	var err error
	ctx := context.TODO()

	err = config.StartConfig()
	FatalError(err)

	db := GetDatabase(ctx)

	err = api.New(db).Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	return database.New(client)
}
