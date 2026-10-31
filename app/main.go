package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Failed to close listener")
			os.Exit(1)
		}
	}(listener)

	fmt.Println("Listening on port 4221")

	for {
		// Block until we receive an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		// Handle client connection
		HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	// Read request
	reqBytes := make([]byte, 1024)
	n, err := conn.Read(reqBytes)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}

	// Whole request
	reqMessage := string(reqBytes[:n])

	// Whole (splitted) request
	reqParts := strings.Split(reqMessage, "\r\n")

	// Handle url
	rl, err := New(reqParts[0]) // Only get the request line
	if err != nil {
		fmt.Println("Error handling url: ", err.Error())
	}

	// response variable
	var res string

	// Check if endpoint empty
	if rl.RequestTarget == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	} else {
		// get endpoint parts
		ep := rl.getParsedTaget()

		switch ep[1] {
		case "index.html":
			res = "HTTP/1.1 200 OK\r\n\r\n"
		case "echo":
			if len(ep) < 3 {
				res = "HTTP/1.1 200 OK\r\n\r\n"
				break
			}
			res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(ep[2]), ep[2])
		default:
			res = "HTTP/1.1 404 Not Found\r\n\r\n"
		}
	}

	fmt.Println("*RECEIVED DATA*:\n", reqMessage)

	// Write response the connection
	_, err = conn.Write([]byte(res))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		return
	}
}
