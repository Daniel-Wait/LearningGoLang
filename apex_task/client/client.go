package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	// ...
	wg := &sync.WaitGroup{}
	tsend := func(wg *sync.WaitGroup) {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			fmt.Println(text)
			_, err := fmt.Fprintf(conn, text)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}

	tread := func(wg *sync.WaitGroup) {
		// Create a buffer to read data into
		buffer := make([]byte, 1024)
		for {
			// Read data from the client
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println(string(buffer[:n]))
		}
		wg.Done()
	}
	// Read and process data from the server
	// ...
	wg.Add(2)
	go tsend(wg)
	go tread(wg)
	wg.Wait()
}
