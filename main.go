package main

import (
	"C"
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func handle_connection(connection net.Conn) {
	defer connection.Close()

	///send some data
	_, _ = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)

	for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received: ", string(buffer[:mLen]))
	}
}

//export luaEntryPoint
func luaEntryPoint() int {
	//establish connection
	listener, _ := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println(err.Error())
		}

		go handle_connection(connection)
	}
}

func main() {}
