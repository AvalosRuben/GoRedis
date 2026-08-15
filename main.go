package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")

	// Create the server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error: The server could not be created: ", err)
		return
	}
	
}