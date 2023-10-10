package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zulkhair/taxi-fares/controller/console"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "console", "console/http")
	flag.Parse()

	// Setup log
	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	log.Logger = zerolog.New(file).With().Timestamp().Logger()

	switch mode {
	case "http":
		startHttp()
	case "console":
		startConsole()
	default:
		log.Fatal().Msgf("Mode '%s' not found", mode)
	}
}

func startHttp() {
	// Todo on the next improvement
	fmt.Println("HTTP mode coming soon, still on development")
	fmt.Println("Switching to console mode")
	startConsole()
}

func startConsole() {
	log.Info().Msgf("Starting console")
	c, err := console.New()
	if err != nil {
		log.Fatal().Msgf("Cannot start console, err : %v", err)
	}
	err = c.StartApp()
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}
}
