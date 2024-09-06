package main

import (
	"fmt"
	"net"
	"time"
)

func broadcastToLAN() {
	port := ":8989"
	broadcastAddr, err := getBoardcastAddr("ens33")
	message := "Manager"

	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Broadcast:", broadcastAddr)
	// Set the broadcast address and port
	//broadcastAddr := "192.168.79.255:8989"
	broadcastAddr = broadcastAddr + port
	

	// Create a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", broadcastAddr)
	if err != nil {
			fmt.Println("Error resolving UDP address:", err)
			return
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
			fmt.Println("Error creating UDP connection:", err)
			return
	}
	defer conn.Close()

	fmt.Println("Broadcasting to", broadcastAddr)

	// Broadcast the message every 30 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
			select {
			case <-ticker.C:
					// Send the broadcast message
					_, err := conn.Write([]byte(message))
					if err != nil {
							fmt.Println("Error sending message:", err)
					} else {
							fmt.Println("Message broadcasted:", message)
					}
			}
	}
}