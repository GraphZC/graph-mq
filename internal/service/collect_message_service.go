package service

import (
	"log"

	"github.com/GraphZC/graph-mq/internal/model"
)


type CollectMessageService struct {
	Chanel chan string
	List *model.QueueLinkedList
	SubService *SubscribeService
}

func NewCollectMessageService(list *model.QueueLinkedList, subService *SubscribeService) *CollectMessageService {
	cms := &CollectMessageService{
		Chanel: make(chan string),
		List: list,
		SubService: subService,
	}
	go handleCollect(cms)

	return cms
}

func handleCollect(cms *CollectMessageService) {
	for {
		message := <-cms.Chanel

		log.Println("Collecting message: ", message)
		msg, topic := cms.List.Dequeue()
		cms.SubService.Publish(topic, msg)
	}
}