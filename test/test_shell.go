package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Define the shell script command
	cmd := exec.Command("bash", "-c", "./action/get_join_token.sh")

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running the script: %v", err)
	}

	// Print the output of the script
	fmt.Printf("Script Output: %s\n", output)
}
