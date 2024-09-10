package main

import (
	"fmt"
	"os/exec"
)

func setExecutablePermissions() error {
	actions_list := []string {
		"./actions/join_cluster.sh",
		"./actions/leave_swarm.sh",
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

func joinSwarmCluster(token string, managerAddr string) error {
	cmd := exec.Command("bash", "-c", "./actions/join_cluster.sh " + token + " " + managerAddr)

	_, err := cmd.Output()
	return err
}

func leaveSwarm() error {
	cmd := exec.Command("bash", "-c", "./actions/leave_swarm.sh")

	_, err := cmd.Output()
	return err
}

// func main() {
// 	err := joinSwarmCluster("SWMTKN-1-2hb55s7157dkx2nqot3cphvkhuxcpkgpsahxpvqq50bu1jfrvz-b4tungf3rxjq5a89ryfzr65m7", "192.168.79.145:2377")
// 	if err != nil {
// 		fmt.Printf("Error executing script: %s\n", err)
// 	}
// }