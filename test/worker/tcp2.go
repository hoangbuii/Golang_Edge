package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "192.168.79.145:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Start reading user input and sending it to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message (type 'exit' to disconnect): ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		// Send the message to the server
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending to server:", err)
			return
		}

		// If the client sends "exit", break the loop to disconnect
		if message == "exit" {
			fmt.Println("Disconnecting from server...")
			break
		}

		// Read the server's response
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Received from server:", response)
	}
}
