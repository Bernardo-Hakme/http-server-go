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
	// Read request
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}

	// response variable (can't define it in if-statement because of block scope)
	var res []byte

	// if request != "GET / HTTP/1.1... --> 404 Not Found
	// else --> 200 OK

	if string(buf[5]) != " " {
		res = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	} else {
		res = []byte("HTTP/1.1 200 OK\r\n\r\n")
	}

	fmt.Println("*RECEIVED DATA*:\n", string(buf))

	// Write response the connection
	_, err = conn.Write(res)
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		return
	}
}
