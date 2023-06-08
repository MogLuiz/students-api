package main

import (
	"log"

	"github.com/MogLuiz/students-api/api"
	"github.com/MogLuiz/students-api/infra/config"
)

func main() {
	var err error

	err = config.StartConfig()
	FatalError(err)

	err = api.New().Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
