package main

import (
	"log"

	"github.com/vvalvvizal/go_study/orchestrator/internal/config"
	"github.com/vvalvvizal/go_study/orchestrator/internal/infra/mysql"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	db, err := mysql.Connect(cfg)
	if err != nil {
		log.Fatalf("mysql: %v", err)
	}
	defer db.Close()

	log.Println("DB connected, port=", cfg.Port)
}
