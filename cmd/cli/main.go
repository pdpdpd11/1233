// Package main is the entry point for the Bepass application.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bepass-org/bepass/config"
	"github.com/bepass-org/bepass/logger"
	"github.com/bepass-org/bepass/server"
)

func main() {
	var (
		configPath string
		verbose    bool
	)
	flag.StringVar(&configPath, "config", "./config.json", "Path to configuration file")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode")
	flag.Parse()

	// Load and validate configuration from the JSON file
	err := loadConfig(configPath)
	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleShutdown(cancel)

	// Run the server with the loaded configuration and context for graceful shutdown
	err = server.Run(verbose)
	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}
}

func loadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config.G)
	if err != nil {
		if strings.Contains(err.Error(), "invalid character") {
			return fmt.Errorf("configuration file is not valid JSON")
		}
		return err
	}

	return nil
}

func handleShutdown(cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	<-c
	fmt.Println("\nShutting down gracefully...")

	// Call the cancel function to trigger the shutdown.
	cancelFunc()

	// Now, you can add additional shutdown logic here if needed before exiting.
	// For example, waiting for server to fully shutdown, closing database connections, etc.

	os.Exit(0)
}
