package main

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Token  string       `yaml:"token" env:"BOT_TOKEN" env-description: "Telegram bot token" env-required: "true"`
	Server ServerConfig `yaml:"server"`
	//GithubToken string       `yaml:"github_token" env:"GITHUB_TOKEN" env-description: "GitHub token" env-required: "true"`
}

type ServerConfig struct {
	Host string `yaml:"host" env:"SERVER_HOST" env-description: "server host" env-required: "true"`
	Port int32  `yaml:"port" env:"SERVER_PORT" env-description: "server port" env-required: "true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {

		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}

	})
	return instance
}
