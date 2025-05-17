package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// Get the project directory from command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <project_directory>")
		os.Exit(1)
	}

	projectDir := os.Args[1]
	log.Printf("dir: %s", projectDir)

	// Check if the provided path is valid
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		log.Fatalf("Error: Directory not found: %s", projectDir)
	} else {
		log.Printf("Directory exist: %s", projectDir)
	}
}
