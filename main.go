package main

import (
	"log"
	"net"

	"github.com/GraphZC/graph-mq/internal/controller"
	"github.com/GraphZC/graph-mq/internal/service"
)

func main() {
	service.InitQueueList()

	// Create the server connection
	listener, err := net.Listen("tcp", ":18060")
	if err != nil {
		log.Println("ERR: ", err)
		return
	}

	// Handle user's message 
	controller.HandleQueue(listener)
}