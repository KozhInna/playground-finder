package main

import (
	"fmt"
	"os"

	"github.com/KozhInna/playground-finder/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(cfg.Port)
}
