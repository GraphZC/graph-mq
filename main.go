package main

import (
	"fmt"
	"net"

	"github.com/GraphZC/mq-socket-programming/internal/service"
)

func main() {
	// Create the server connection
	connection, err := net.Listen("tcp", ":18060")
	if err != nil {
		fmt.Println(err)
		return
	}

	service.InitQueueList()

	for {
		client, err := connection.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleClient(client)
	}
}

func handleClient(client net.Conn) {
	defer client.Close()

	for {
		buf := make([]byte, 1024)
		n, err := client.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
