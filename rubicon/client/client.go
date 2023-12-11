package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	var user = os.Args[1]
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	// ...

	for {
		// Send the message to the server
		message := string("Hello, Server! I am User") + user
		fmt.Fprintf(conn, message)
		time.Sleep(time.Second)
	}

	// Read and process data from the server
	// ...
}
