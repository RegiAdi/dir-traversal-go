package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	log.Printf("Starting project traversal in: %s", projectDir)

	// Define the walk function
	walkFunction := func(path string, d fs.DirEntry, err error) error {
		// --- 1. Handle errors that occurred before calling this function ---
		if err != nil {
			// Log the error and decide if you want to stop or try to continue.
			// Returning the error stops the entire walk.
			// Returning nil attempts to continue past the error.
			// Let's log and skip this specific path/error for this example.
			log.Printf("Error accessing path %s: %v", path, err)
			return nil
		}

		// --- 2. Decide what to do with directories ---
		if d.IsDir() {
			name := d.Name()
			// Skip common version control, dependency, or build output directories
			if name == ".git" ||
				name == "vendor" ||
				name == "node_modules" ||
				name == "target" || // Common Rust build dir
				name == "dist" ||
				strings.HasPrefix(name, ".") { // Skip other hidden dirs
				log.Printf("Skipping directory and its contents: %s", path)
				return filepath.SkipDir // Tell WalkDir to not go into this directory
			}
		} else {
			// --- 3. Decide what to do with files ---
			// Check file extension or name to identify source files
			fmt.Println(path)

			if strings.HasSuffix(d.Name(), ".go") {
				// Found a Go source file!
				// fmt.Println(path) // Print the path

				// --- Optional: Read the file content ---
				// content, readErr := os.ReadFile(path)
				// if readErr != nil {
				// 	log.Printf("Error reading file %s: %v", path, readErr)
				// 	// Decide if this read error should stop the whole walk (return readErr)
				// 	// or just skip processing this file (return nil).
				// } else {
				// 	// --- Process the file content ---
				// 	// This is where you would integrate with a parser like Tree-sitter
				// 	// Example: Pass 'content' to your Tree-sitter parser
				// 	// parseTree := yourParser.Parse(content, nil)
				// 	// ... then traverse the parseTree ...
				// }
			}
			// You could add checks for other file types like ".rs", ".py", ".js", etc.
		}

		// --- 4. Return nil to continue the walk ---
		return nil
	}

	// Call WalkDir with the root path and the function
	err := filepath.WalkDir(projectDir, walkFunction)

	// Handle any error that caused WalkDir to stop prematurely
	if err != nil {
		log.Fatalf("Error during directory traversal: %v", err)
	}

	log.Println("Project traversal finished.")
}
