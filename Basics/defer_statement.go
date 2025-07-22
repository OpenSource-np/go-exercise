package basics

import (
	"fmt"
	"os"
)

// Defer Statement
// Learn how to utilize the defer statement in Go for resource management and cleanup tasks.

// BasicDefer demonstrates the use of defer statement to ensure that a file is closed after writing.
func BasicDefer() {
	f, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("File content")
	if err != nil {
		panic(err)
	}

	fmt.Println("File written successfully")
}

// DeferMultiple demonstrates the use of multiple defer statements.
func DeferMultiple() {
	defer fmt.Println("Bob")
	defer fmt.Println("Alex")
	fmt.Println("Welcome")
}

// DeferInFunc demonstrates the use of defer in a function to clean up resources.
// It ensures that cleanup is called after the function execution, even if an error occurs.
func DeferInFunc() {
	defer cleanup()
	fmt.Println("Processing...")
}

func cleanup() {
	fmt.Println("Cleaning up resources.")
}
