package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server at IP 192.168.1.10 and port 8080
	// Replace with the server IP if the server is on another device
	serverAddress := "192.168.1.10:8080"
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server:", serverAddress)


	for {
		var message string
		fmt.Scanln(&message)
		// Send a "connect" signal to the server
		//message := "connect"
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Printf("Sent message to server: %s\n", message)
	}
}
