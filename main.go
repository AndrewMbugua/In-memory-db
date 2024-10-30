package main

import (
	"fmt"
	"net"
)

// Server - Mini Redis-like server
func Server() {
	fmt.Println("Listening on port :6379")

	server, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	connection, err := server.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	for {

		resp := NewResp(connection)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)

		// Ignore the request and send below string
		connection.Write([]byte("+Server Responds, OK\r\n"))

	}
}

func main() {

	//Server()
	fmt.Println("Listening on port :6379")

	server, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	connection, err := server.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	for {

		resp := NewResp(connection)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)

		// Ignore the request and send below string
		connection.Write([]byte("+Server Responds, OK\r\n"))

	}

}
