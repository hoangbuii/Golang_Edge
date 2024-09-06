package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen on TCP port 8080 on any interface
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error setting up server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is waiting for client connection on port 8080...")

	// Wait for a connection from the client
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Client connected from", conn.RemoteAddr())
	for {
		// Read data sent by the client
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
	

		// Display the message sent by the client
		fmt.Printf("Received message: %s\n", string(buffer[:n]))
	}
}
