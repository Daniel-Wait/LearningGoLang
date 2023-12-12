package main

import (
	"fmt"
	"net"
)

var client_map map[string]net.Conn

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	// Create Clients map
	client_map = make(map[string]net.Conn)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Show who is connecting
	tcp_addr := fmt.Sprintf("%s", conn.RemoteAddr())
	fmt.Printf("Connection from: %s\n", tcp_addr)
	_, found := client_map[tcp_addr]
	if found == false {
		client_map[tcp_addr] = conn
	}
	defer delete(client_map, tcp_addr)

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data
		newmessage := fmt.Sprintf("%s: %s", conn.RemoteAddr(), buffer[:n])
		fmt.Printf(newmessage)
		for _, map_conn := range client_map {
			// if _ != tcp_addr {
			// }
			map_conn.Write([]byte(newmessage))
		}
	}
}
