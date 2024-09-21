package main

import (
	"fmt"
	"log"
)

func main() {
	err := setExecutablePermissions()
	if err != nil {
		log.Fatalf("Error set excutable for the script: %v", err)
	}
	go broadcastToLAN("ens33", 8989, "Manager")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "q" || input == "Q" {
			fmt.Println("Bye!")
			break
		}
	}
}
