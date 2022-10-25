package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Token    string         `yaml:"token" env:"BOT_TOKEN" env-description: "Telegram bot token" env-required: "true"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	//GithubToken string       `yaml:"github_token" env:"GITHUB_TOKEN" env-description: "GitHub token" env-required: "true"`
}

type ServerConfig struct {
	Host string `yaml:"host" env:"SERVER_HOST" env-description: "server host" env-required: "true"`
	Port int32  `yaml:"port" env:"SERVER_PORT" env-description: "server port" env-required: "true"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" env-description: "database host" env-required: "true"`
	Port     int32  `yaml:"port" env:"DATABASE_PORT" env-description: "database port" env-required: "true"`
	User     string `yaml:"user" env:"DATABASE_USER" env-description: "database user" env-required: "true"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-description: "database password" env-required: "true"`
	Instance string `yaml:"instance" env:"DATABASE_INSTANCE" env-description: "database instance" env-required: "true"`
}

const configPath = "configs/config.yaml"

var singleton *Config
var load sync.Once

func GetConfig() *Config {
	load.Do(func() {
		singleton = &Config{}
		if err := cleanenv.ReadConfig(configPath, singleton); err != nil {
			help, _ := cleanenv.GetDescription(singleton, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return singleton
}
