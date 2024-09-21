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

		parts := strings.Split(message, "|")
		command := parts[0]

		if command == "SCAN" {
			token, err := getJoinToken()
			if err != nil {
				log.Fatalf("Error getting join token: %v", err)
			}
			token = strings.TrimSuffix(token, "\n")
			information = "INFO|" + token + "|2377" 
			_, err = conn.WriteToUDP([]byte(information), addr)
			if err != nil {
				fmt.Println("Error sending echo message:", err)
			} else {
				fmt.Println("send message back to", addr)
			}
		}

		fmt.Printf("Received message from %s: %s\n", addr, message)

		
	}
}