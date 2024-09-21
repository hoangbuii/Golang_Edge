package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// Calculate the broadcast address
func calculateBroadcast(ip net.IP, mask net.IPMask) net.IP {
	ip = ip.To4()
	broadcast := make(net.IP, len(ip))
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}
	return broadcast
}

// Get the default gateway (Unix-based systems)
func getDefaultGateway(iface string) (string, error) {
	out, err := exec.Command("ip", "route", "show", "default", "0.0.0.0/0").Output()
	if err != nil {
		return "", err
	}
	parts := strings.Fields(string(out))
	if len(parts) > 2 && parts[0] == "default" && parts[1] == "via" {
		return parts[2], nil
	}
	return "", fmt.Errorf("default gateway not found")
}

// Get the DNS servers from /etc/resolv.conf (Unix-based systems)
func getDNSServers() ([]string, error) {
	out, err := exec.Command("cat", "/etc/resolv.conf").Output()
	if err != nil {
		return nil, err
	}

	var dnsServers []string
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				dnsServers = append(dnsServers, parts[1])
			}
		}
	}
	return dnsServers, nil
}

func showNetworkConfiguration() {
	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		// Skip down or loopback interfaces
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		fmt.Println("Interface Name:", iface.Name)

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error fetching addresses:", err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}

			if ipNet.IP.To4() != nil {
				fmt.Println("  IP Address:", ipNet.IP.String())
				fmt.Println("  Subnet Mask:", ipNet.Mask.String())

				// Calculate and print broadcast address
				broadcast := calculateBroadcast(ipNet.IP, ipNet.Mask)
				fmt.Println("  Broadcast Address:", broadcast.String())
			}
		}

		// Fetch default gateway
		gateway, err := getDefaultGateway(iface.Name)
		if err == nil {
			fmt.Println("  Default Gateway:", gateway)
		} else {
			fmt.Println("  Default Gateway: Not found")
		}

		// Fetch DNS servers
		dnsServers, err := getDNSServers()
		if err == nil {
			fmt.Println("  DNS Servers:", dnsServers)
		} else {
			fmt.Println("  DNS Servers: Not found")
		}

		fmt.Println()
	}
}

func getBoardcastAddr(interfaceName string) (string, error) {
	// Get the network interface by name
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", fmt.Errorf("error getting interface %s: %v", interfaceName, err)
	}

	// Get the addresses associated with the interface
	addrs, err := iface.Addrs()
	if err != nil {
		return "", fmt.Errorf("error getting addresses for interface %s: %v", interfaceName, err)
	}

	// Loop through the addresses to find the first IPv4 address
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.To4() == nil {
			continue
		}

		// Calculate the broadcast address
		broadcast := calculateBroadcast(ipNet.IP, ipNet.Mask)
		return broadcast.String(), nil
	}

	return "", fmt.Errorf("no IPv4 address found for interface %s", interfaceName)
}