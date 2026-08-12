package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Host        string        `yaml:"host"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Storage struct {
	Path     string `yaml:"path"`
	Provider string `yaml:"provider"`
	Url      string `yaml:"url"`
}

type Config struct {
	AppName     string `yaml:"app_name"`
	Env         string `yaml:"env"`
	LoggerLevel string `yaml:"logger_level"`
	HttpServer  `yaml:"http_server"`
	Storage     `yaml:"storage"`
}

func LoadConfig() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "config/local.yaml", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config file path not loaded")
		}
	}

	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatalf("config file doesn't exist at path: %v", err)
	}

	var cfg Config
	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("config file not accroding to requirements")
	}
	return &cfg
}
