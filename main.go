package main

import (
	"log"
	"net"

	"github.com/GraphZC/mq-socket-programming/internal/controller"
	"github.com/GraphZC/mq-socket-programming/internal/service"
)

func main() {
	service.InitQueueList()

	// Create the server connection
	listener, err := net.Listen("tcp", ":18060")
	if err != nil {
		log.Println(err)
		return
	}

	// Handle user's message 
	controller.HandleQueue(listener)
}