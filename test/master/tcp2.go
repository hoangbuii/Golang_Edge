package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	clients   = make(map[string]net.Conn)
	clientsMu sync.Mutex // To synchronize access to the clients map
)

func main() {
	// Start the server and listen for connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	// Goroutine to handle sending messages to specific clients from the server
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter client address and message (e.g., 127.0.0.1:12345 Hello): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Split the input into address and message parts
			parts := strings.SplitN(input, " ", 2)
			if len(parts) < 2 {
				fmt.Println("Invalid input. Usage: <address> <message>")
				continue
			}

			address, message := parts[0], parts[1]

			// Send the message to the specified client
			clientsMu.Lock()
			conn, ok := clients[address]
			clientsMu.Unlock()

			if !ok {
				fmt.Println("Client not found:", address)
				continue
			}

			_, err := conn.Write([]byte("Server: " + message + "\n"))
			if err != nil {
				fmt.Println("Error sending message to client:", err)
				continue
			}

			fmt.Println("Message sent to", address)
		}
	}()

	for {
		// Accept new client connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		clientAddr := conn.RemoteAddr().String()
		fmt.Println("Client connected:", clientAddr)

		// Store the client connection in the map
		clientsMu.Lock()
		clients[clientAddr] = conn
		clientsMu.Unlock()

		// Handle the client in a separate goroutine
		go handleClient(conn, clientAddr)
	}
}

func handleClient(conn net.Conn, clientAddr string) {
	defer func() {
		clientsMu.Lock()
		delete(clients, clientAddr)
		clientsMu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", clientAddr)
	}()

	reader := bufio.NewReader(conn)
	for {
		// Read the message from the client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		message = strings.TrimSpace(message)

		// Check if the client wants to disconnect
		if message == "exit" {
			fmt.Println("Client requested to disconnect:", clientAddr)
			return // Gracefully close the connection
		}

		fmt.Println("Received from", clientAddr+":", message)

		// Echo the message back to the client
		_, err = conn.Write([]byte("Server echo: " + message + "\n"))
		if err != nil {
			fmt.Println("Error sending to client:", err)
			return
		}
	}
}
