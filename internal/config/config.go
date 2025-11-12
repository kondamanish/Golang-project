package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string                    `yaml:"env" env-required:"true"`          // environment variable  struct tags
	StoragePath string                    `yaml:"storage_path" env-required:"true"` // storage path for the database annotations for the yaml file
	HTTPServer  `yaml:"http_server_port"` // http server configuration annotations for the yaml file
}

func MustLoad() *Config {
	var configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file default is config/local.yaml")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatalf("config path is required: %v", configPath)
		}

	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) { //this is to check if the config file exists
		log.Fatalf("config file does not exist: %v", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg) // error handling for the config file  // this is to read the config file
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}
