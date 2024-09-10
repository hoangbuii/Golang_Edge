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
	
	// state [idle, connected, disconnected]
	state := "idle"

	reader := bufio.NewReader(conn)
	for {
		// Read the message from the client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		message = strings.TrimSpace(message)

		parts := strings.Split(message, "|")

		command := parts[0]

		if command == "connect" && (state == "idle" || state == "disconnected") {
			join_token := "dbafkhebbdsahfjrhfrYBSh7DBGEH&hndjf"
			_, err = conn.Write([]byte("Server echo: " + join_token + "\n"))
			if err != nil {
				fmt.Println("Error sending to client:", err)
				return
			}
			state = "connected"
		}

		if command == "done" && (state == "connected") {
			fmt.Println("ok")
		}

		if command == "exit" {
			fmt.Println("Client requested to disconnect:", clientAddr)
			state = "disconnected"
			return // Gracefully close the connection
		}

		
		//handleMessage(message)

		// Check if the client wants to disconnect
		// if message == "exit" {
		// 	fmt.Println("Client requested to disconnect:", clientAddr)
		// 	return // Gracefully close the connection
		// }

		fmt.Println("Received from", clientAddr+":", message)

		// Echo the message back to the client
		_, err = conn.Write([]byte("Server echo: " + message + "\n"))
		if err != nil {
			fmt.Println("Error sending to client:", err)
			return
		}
	}
}

// func handleMessage(message string) {
// 	parts := strings.Split(message, "|")

// 	command := parts[0]

// 	if command == "connect" 
// }