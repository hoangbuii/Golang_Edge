package main

import (
	"fmt"
	"log"
	"os/exec"
)

func setExecutablePermissions() error {
	actions_list = []strings {
		"./actions/get_join_token.sh"
	}

	for index, action := range actions_list {
		cmd := exec.Command("bash", "-c", "chmod +x " + action)
		
		output, err := cmd.Output()
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