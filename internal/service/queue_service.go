package service

import "github.com/GraphZC/graph-mq/internal/model"

var queue *model.QueueLinkedList

func InitQueueList() {
	queue = model.NewQueueLinkedList()
}

func EnqueueMessage(topic string, message string) {
	queue.Enqueue(topic, message)
}

func DequeueMessage() (message string, topic string) {
	return queue.Dequeue()
}

func PrintQueue() {
	queue.PrintQueue()
}

func GetQueue() *model.QueueLinkedList {
	return queue
}


