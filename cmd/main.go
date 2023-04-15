package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jvhab/Redis-crud/config"
	"log"
)

func main() {
	cfg := &config.Config{}
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		log.Fatalf("dont read config")
	}
	fmt.Println(cfg)
}
