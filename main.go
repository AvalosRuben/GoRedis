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

	conn, err := l.Accept()
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}

	defer conn.Close()


	for{
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil{
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}