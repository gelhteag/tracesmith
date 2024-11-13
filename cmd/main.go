package main

import (
	"fmt"
	"os"

	"github.com/gelhteag/tracesmith/pkg/tracesmith"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Warning: Error loading .env file: %v", err)
	}

	// Retrieve required environment variables
	langchainEndpoint := os.Getenv("LANGCHAIN_ENDPOINT")
	langchainAPIKey := os.Getenv("LANGCHAIN_API_KEY")
	langchainProject := os.Getenv("LANGCHAIN_PROJECT")

	if langchainEndpoint == "" || langchainAPIKey == "" || langchainProject == "" {
		log.Fatal("Missing one or more required environment variables: LANGCHAIN_ENDPOINT, LANGCHAIN_API_KEY, LANGCHAIN_PROJECT")
	}

	// Set environment variables for internal use
	os.Setenv("LANGCHAIN_ENDPOINT", langchainEndpoint)
	os.Setenv("LANGCHAIN_API_KEY", langchainAPIKey)
	os.Setenv("LANGCHAIN_PROJECT", langchainProject)

	client := tracesmith.NewClient()
	chain := tracesmith.NewChain(client, "Embedding")

	// Start a run
	inputs := map[string]interface{}{"query": "example input"}
	run, err := chain.AddRun("Embedding", "chain", inputs, nil)
	if err != nil {
		log.WithError(err).Fatal("Error starting run")
	}

	// Example processing
	fmt.Println("Processing...")

	// End the run
	outputs := map[string]interface{}{"result": "example output"}
	if err := run.End(outputs); err != nil {
		log.WithError(err).Fatal("Error ending run")
	}

	// End all runs in the chain (optional)
	if err := chain.EndAllRuns(outputs); err != nil {
		log.WithError(err).Fatal("Error ending all runs in chain")
	}

	log.Infoln("Run completed successfully")
}
