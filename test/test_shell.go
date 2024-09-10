package main

import (
	"fmt"
	"log"
	"os/exec"
)

func setExecutablePermissions() error {
	actions_list := []string {
		"./actions/get_join_token.sh",
		"./actions/list_node.sh",
	}

	for _, action := range actions_list {
		cmd := exec.Command("bash", "-c", "chmod +x " + action)
		
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing script: %s\n", err)
			return err
		}
	}
	return nil
}

func getJoinToken() (string, error) {
	// Define the shell script command
	cmd := exec.Command("bash", "-c", "./actions/get_join_token.sh")

	// Run the command and capture the output
	output, err := cmd.Output()
	return output, err
	
}

func main() {
	err := setExecutablePermissions()
	if err != nil {
		log.Fatalf("Error set excutable for the script: %v", err)
	}
	token, err := getJoinToken()

	if err != nil {
		log.Fatalf("Error running the script: %v", err)
	}
	fmt.Printf("Script Output: %s\n", token)
	
}
