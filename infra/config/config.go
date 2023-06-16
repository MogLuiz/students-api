package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT          int    `envconfig:"PORT"`
	MongoURL      string `envconfig:"MONGO_URL"`
	MongoDatabase string `envconfig:"MONGO_DATABASE"`
}

var Env Config

func StartConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
