package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BenjaminAHawker/hawk-bot/internal/bot"
	"github.com/BenjaminAHawker/hawk-bot/internal/config"
	"github.com/BenjaminAHawker/hawk-bot/internal/db"
)

func main() {
	cfg := config.Load()
	log.Println("Config loaded")

	pgPool := db.Connect(cfg)
	defer pgPool.Close()

	// Run migrations on startup
	if err := pgPool.Migrate(context.Background()); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	dg, err := bot.NewSession(cfg.DiscordToken, pgPool)
	if err != nil {
		log.Fatalf("Failed to create Discord session: %v", err)
	}
	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	log.Println("Shutting down gracefully...")
}
