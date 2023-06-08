package main

import (
	"log"

	"github.com/MogLuiz/students-api/api"
)

func main() {
	if err := api.New().Start(); err != nil {
		log.Fatal(err)
	}
}
