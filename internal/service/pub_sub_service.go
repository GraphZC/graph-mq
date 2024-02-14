package service

import (
	"log"
	"net"
)

type SubscribeTopic struct {
	Subscribers []net.Conn
	Chanel chan string
}

type SubscribeService struct {
	Subscribes map[string]*SubscribeTopic
}

func NewSubscribeService() *SubscribeService {
	ss := &SubscribeService{
		Subscribes: make(map[string]*SubscribeTopic),
	}
	go handlePublish(ss)

	return ss
}

func (ss *SubscribeService) Subscribe(topic string, client net.Conn) {
	_, ok := ss.Subscribes[topic]
	if !ok {
		ss.Subscribes[topic] = &SubscribeTopic{
			Subscribers: []net.Conn{client},
			Chanel: make(chan string),
		}
	} else {
		ss.Subscribes[topic].Subscribers = append(ss.Subscribes[topic].Subscribers, client)
	}
}

func (ss *SubscribeService) Publish(topic string, message string) {
	sub, ok := ss.Subscribes[topic]
	if !ok {
		return
	}
	
	log.Println("Publishing to topic: ", topic)

	sub.Chanel <- message
}

func handlePublish(ss *SubscribeService) {
	for {
		for _, sub := range ss.Subscribes {
			msg := <-sub.Chanel
			for _, client := range sub.Subscribers {
				client.Write([]byte(msg))
			}
		}
	}
}