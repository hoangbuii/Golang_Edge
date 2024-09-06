package main

import (
	"fmt"
	"net"
)

func listenBroadcast() {
	// Set the port to listen on (same as the broadcast port)
	listenAddr := ":8989"

	// Create a UDP address for listening
	udpAddr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection for listening
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening for broadcast messages on port 8989")

	buffer := make([]byte, 1024)

	for {
		// Read incoming UDP messages
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			continue
		}

		// Print the received message
		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", addr, message)
	}
}