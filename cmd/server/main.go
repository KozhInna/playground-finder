package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/KozhInna/playground-finder/internal/config"
	"github.com/KozhInna/playground-finder/internal/db"
	"github.com/gin-gonic/gin"
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

	ctx := context.Background()

	// Open the connection pool and ping it
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer pool.Close()

	// Router with the default logger and panic recovery middleware
	r := gin.Default()

	// Liveness check of the server
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Start the server (blocks until it stops)
	if err := r.Run(":" + cfg.Port); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
