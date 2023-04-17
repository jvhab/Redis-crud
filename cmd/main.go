package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jvhab/Redis-crud/config"
	"github.com/jvhab/Redis-crud/internal/server"
	"log"
)

func main() {
	cfg := &config.Config{}
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		log.Fatalf("dont read config")
	}
	server.Run(cfg)
}
