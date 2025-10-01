package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <moduleName>")
		os.Exit(1)
	}

	moduleName := os.Args[1]
	logFile := fmt.Sprintf("result/%s.log", moduleName)

	// Ensure the result directory exists
	if err := os.MkdirAll("result", 0755); err != nil {
		fmt.Printf("Error creating result directory: %v\n", err)
		os.Exit(1)
	}

	// Open log file in append mode
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	cmd := exec.Command("go", "run", "cmd/crawler/main.go", moduleName)

	// Redirect stdout and stderr to the log file
	cmd.Stdout = file
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}

}
