package controller

import (
	"log"
	"net"

	"github.com/GraphZC/graph-mq/internal/service"
)

func HandleQueue(listener net.Listener) {
	subscribeService := service.NewSubscribeService()
	collectMessageService := service.NewCollectMessageService(service.GetQueue(), subscribeService)

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go handleClient(client, collectMessageService)
	}
}

func handleClient(client net.Conn, collectMessageService *service.CollectMessageService) {
	defer client.Close()

	for {
		buf := make([]byte, 1024)

		// COMMAND: ARG1, ARG2
		size, err := client.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		msg := string(buf[:size])

		cmd := service.ExtractCommand(msg)

		switch cmd.Cmd {
		case "PUBL":
			// PUBL:TOPIC,MESSAGE
			if len(cmd.Arguments) != 2 {
				log.Println("401 BAD REQUEST")
				return
			}
			service.EnqueueMessage(cmd.Arguments[0], cmd.Arguments[1])
			collectMessageService.Chanel <- cmd.Arguments[0]

			client.Write([]byte("201 PUBLISHED\n"))
		case "SUBC":
			// SUBC:TOPIC
			if len(cmd.Arguments) != 1 {
				log.Println("401 BAD REQUEST")
				return
			}
			collectMessageService.SubService.Subscribe(cmd.Arguments[0], client)
			client.Write([]byte("200 SUBSCRIBED\n"))

		default:
			log.Println("400 INVALID COMMAND")
		}

	}
}
