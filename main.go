package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zulkhair/taxi-fares/controller/console"
	"github.com/zulkhair/taxi-fares/domain/config"
	"gopkg.in/yaml.v2"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "console", "console/http")
	flag.Parse()

	// Read Config file
	c := readConfigFile()
	if c == nil {
		log.Info().Msgf("Using default values")
	}

	// Setup log
	logFile := "log/app.log"
	if c != nil && c.Log.File != "" {
		logFile = c.Log.File
	}
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(logFile), 0700)
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal().Msgf("Cannot open log file, err : %v", err)
	}
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

func readConfigFile() *config.Config {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Error().Msgf("Cannot open config file, err : %v", err)
		return nil
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Error().Msgf("Cannot read config file, err : %v", err)
		return nil
	}

	var cfg config.Config
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		log.Error().Msgf("Cannot unmarshal config file, err : %v", err)
		return nil
	}
	return &cfg
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
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
}
