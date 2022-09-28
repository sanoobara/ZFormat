package main

import (
	"log"
	"rm/config/cleanenv"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	//app.Run(cfg)
}
