package service

import "github.com/GraphZC/mq-socket-programming/internal/model"

var Queue *model.QueueLinkedList

func InitQueueList() {
	Queue = model.NewQueueLinkedList()
}
