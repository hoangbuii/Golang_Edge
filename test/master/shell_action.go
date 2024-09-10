package main

import (
	"fmt"
	"os/exec"
)

func setExecutablePermissions() error {
	actions_list := []string {
		"./actions/get_join_token.sh",
		"./actions/list_node.sh",
		"./actions/remove_down_node.sh",
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
	token := string(output)
	return token, err
	
}

func listSwarmNode() {
	cmd := exec.Command("bash", "-c", "./actions/list_node.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing script: %s\n", err)
		return
	}
	fmt.Printf("Node list:\n%s\n", output)
}

func removeDownNode() {
	cmd := exec.Command("bash", "-c", "./actions/remove_down_node.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing script: %s\n", err)
		return
	}
	fmt.Printf("Node(s) had been removed:\n%s\n", output)
}