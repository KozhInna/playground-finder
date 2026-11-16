package main

import (
	"context"
	"fmt"
	"os"

	"github.com/KozhInna/playground-finder/internal/config"
	"github.com/KozhInna/playground-finder/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env into the environment (ignored if the file is missing)
	godotenv.Load()

	// Read and validate PORT and DATABASE_URL
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(cfg.Port)

	ctx := context.Background()

	// Open the connection pool and ping it
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer pool.Close()
	
	// Check the DB works: count rows in playgrounds
 	var count int
	err = pool.QueryRow(ctx, "SELECT count(*) FROM playgrounds").Scan(&count)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("playgrounds:", count)
}
