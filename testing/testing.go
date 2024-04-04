// Package testing provides utility functions to run unit and integration tests.
package testing

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	RunUnitTests()
	RunIntegrationTests()
}

// RunUnitTests runs all unit tests.
func RunUnitTests() {
	// Command to run `go test -v ./...`
	cmd := exec.Command("go", "test", "-v", "./...")

	// Set stdout and stderr to os.Stdout and os.Stderr respectively
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running go test:", err)
		return
	}
	fmt.Println("Tests executed successfully")
}

// RunIntegrationTests runs all integration tests.
func RunIntegrationTests() {
	// Command to run `go test -v -tags=integration ./...`
	cmd := exec.Command("go", "test", "-v", "-tags=integration", "./...")

	// Set stdout and stderr to os.Stdout and os.Stderr respectively
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running go test:", err)
		return
	}
	fmt.Println("Integration tests executed successfully")
}
