package main

import (
	"fmt"
	"log"

	"github.com/zyaeger/blog-agg/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error finding config: %v", err)
	}
	fmt.Printf("Read Config: %+v\n", cfg)

	err = cfg.SetUser("zach")
	if err != nil {
		log.Fatalf("error setting current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}