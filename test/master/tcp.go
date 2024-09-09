// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	// Listen on TCP port 8080 on any interface
// 	listener, err := net.Listen("tcp", ":8080")

// 	if err != nil {
// 		fmt.Println("Error setting up server:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server is waiting for client connection on port 8080...")

// 	// Wait for a connection from the client
// 	conn, err := listener.Accept()
// 	if err != nil {
// 		fmt.Println("Error accepting connection:", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()

// 	fmt.Println("Client connected from", conn.RemoteAddr())
// 	for {
// 		// Read data sent by the client
// 		buffer := make([]byte, 1024)
// 		n, err := conn.Read(buffer)
// 		if err != nil {
// 			fmt.Println("Error reading from connection:", err)
// 			return
// 		}
	

// 		// Display the message sent by the client
// 		fmt.Printf("Received message: %s\n", string(buffer[:n]))
// 	}
// }


package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle the client in a separate goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Read the message from the client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		fmt.Println("Received from client:", message)

		// Echo the message back to the client
		_, err = conn.Write([]byte("Server: " + message))
		if err != nil {
			fmt.Println("Error sending to client:", err)
			return
		}
	}
}
