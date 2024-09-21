package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

func broadcastToLAN(iface string, port int, managerID string) {
	// port := 8989
	broadcastAddr, err := getBoardcastAddr(iface)
	message := managerID
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Broadcast:", broadcastAddr)
	// Set the broadcast address and port
	//broadcastAddr := "192.168.79.255:8989"
	broadcastAddr = broadcastAddr + ":" +  strconv.Itoa(port)
	

	// Create a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", broadcastAddr)
	if err != nil {
			fmt.Println("Error resolving UDP address:", err)
			return
	}

	// Create a UDP connection to broadcast
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
			fmt.Println("Error creating UDP connection:", err)
			return
	}
	defer conn.Close()

	// Create a UDP connection to listen message
	listenConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error creating UDP listener:", err)
		return
	}
	defer listenConn.Close()

	done := make(chan bool)

	go func() {
		buffer := make([]byte, 1024)
		for {
			n, remoteAddr, err := listenConn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error receiving echo message:", err)
				continue
			}
			fmt.Println("Received echo from", remoteAddr, ":", string(buffer[:n]))
		}
	}()

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

	<-done
}

func scanManager() {

}