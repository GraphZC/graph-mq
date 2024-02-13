package controller

import (
	"log"
	"net"

	"github.com/GraphZC/mq-socket-programming/internal/service"
)

func HandleQueue(listener net.Listener) {
	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go handleClient(client)
	}
}

func handleClient(client net.Conn) {
	defer client.Close()

	for {
		buf := make([]byte, 1024)

		// SEND:TOPIC, MESSAGE
		size, err := client.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		msg := string(buf[:size])

		cmd := service.ExtractCommand(msg)

		switch cmd.Cmd {
		case "SEND":
			if len(cmd.Arguments) != 2 {
				log.Println("400 BAD REQUEST")
				return
			}
			service.EnqueueMessage(cmd.Arguments[0], cmd.Arguments[1])
			service.PrintQueue()
		}
	}
}
