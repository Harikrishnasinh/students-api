package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string
}

type Config struct {
	Env         string     `yaml:"env" env-default:"production" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	// Implementation for loading configuration goes here
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "Path to configuration file")
		flag.Parse()
		fmt.Println("chk", *flags)
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist at path: %s", configPath)
	}

	var cfg Config
	// Load configuration from file and environment variables into cfg
	// This is a placeholder for actual loading logic

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	log.Printf("Configuration loaded successfully: %+v", cfg)

	return &cfg
}
