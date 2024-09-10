package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"log"
)

func setupTCPConnection(managerIP string, port int) {
	// Connect to the server
	conn, err := net.Dial("tcp", managerIP + ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Start reading user input and sending it to the server
	//reader := bufio.NewReader(os.Stdin)

	message := "connect|Manager|12345678"
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("Error sending to server:", err)
		return
	}
	fmt.Println("Sent connect command to server")
	// state [idle, connecting, connected, disconnected]
	state := "connecting"
	for {
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Received from server:", response)
		// response = strings.TrimSpace(response)
		
		parts := strings.Split(response, "|")
		command := parts[0]

		if command == "info" && state == "connecting" {
			token := parts[2]
			managerAddr := managerIP + ":" + parts[1]
			err := joinSwarmCluster(token, managerAddr)
			if err != nil {
				fmt.Println("Erorr to join cluster:", err)
				return
			}
			message = "done|0"
			_, err = conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending to server:", err)
				return
			}
			state = "connected"
		}


	}
	



	// for {
	// 	fmt.Print("Enter message (type 'exit' to disconnect): ")
	// 	message, _ := reader.ReadString('\n')
	// 	message = strings.TrimSpace(message)

	// 	// Send the message to the server
	// 	_, err = conn.Write([]byte(message + "\n"))
	// 	if err != nil {
	// 		fmt.Println("Error sending to server:", err)
	// 		return
	// 	}

	// 	// If the client sends "exit", break the loop to disconnect
	// 	if message == "exit" {
	// 		fmt.Println("Disconnecting from server...")
	// 		break
	// 	}

	// 	// Read the server's response
	// 	response, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Error reading from server:", err)
	// 		return
	// 	}
	// 	fmt.Println("Received from server:", response)
	// }
}

func main() {
	err := setExecutablePermissions()
	if err != nil {
		log.Fatalf("Error set excutable for the script: %v", err)
	}
	setupTCPConnection("192.168.79.145", 8080)
}
