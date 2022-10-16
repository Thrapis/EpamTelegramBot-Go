package main

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-description: "server host" env-required: "true"`
	Port     int32  `yaml:"port" env:"PORT" env-description: "server port" env-required: "true"`
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" env-description: "database host" env-required: "true"`
	Port     int32  `yaml:"port" env:"DATABASE_PORT" env-description: "database port" env-required: "true"`
	User     string `yaml:"user" env:"DATABASE_USER" env-description: "database user" env-required: "true"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-description: "database password" env-required: "true"`
	Instance string `yaml:"instance" env:"DATABASE_INSTANCE" env-description: "database instance" env-required: "true"`
}

var singleton *Config
var load sync.Once

func GetConfig() *Config {
	load.Do(func() {
		singleton = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", singleton); err != nil {
			help, _ := cleanenv.GetDescription(singleton, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return singleton
}
