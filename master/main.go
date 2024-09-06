package main

import (
	"fmt"
)

func main() {
	go broadcastToLAN()
	for {
		var input string
		fmt.Scanln(&input)
		if input == "q" || input == "Q" {
			fmt.Println("Bye!")
			break
		}
	}
}