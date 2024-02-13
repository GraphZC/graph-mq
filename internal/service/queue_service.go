package service

import "github.com/GraphZC/mq-socket-programming/internal/model"

var queue *model.QueueLinkedList

func InitQueueList() {
	queue = model.NewQueueLinkedList()
}

func EnqueueMessage(topic string, message string) {
	queue.Enqueue(topic, message)
}

func PrintQueue() {
	queue.PrintQueue()
}


