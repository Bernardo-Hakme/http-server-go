package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Starts a TCP server listening on localhost:4221
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	// Close server when program finishes
	defer listener.Close()

	fmt.Println("Listening on port 4221")

	for {
		// Block until we receive an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		// Handle client connection
		handleClient(conn)
	}
}

/**
 * Handle the message exchange between client and server
 */
func handleClient(conn net.Conn) {
	// Ensure we close after program ends
	defer conn.Close()

	// Read data
	buf := make([]byte, 1024) // Create slice of bytes to store message
	n, err := conn.Read(buf)  // Store message and return size of slice
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}

	// Print out read data
	fmt.Println("Received data: ", buf[:n])
}
