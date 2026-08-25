package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on port :6379")

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

		_ = value
		writer := NewWriter(conn)
		writer.Write(Value{typ:"string", str: "OK"})

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}