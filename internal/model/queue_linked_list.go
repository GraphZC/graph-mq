package model

import "fmt"

type Node struct {
	Message string
	Topic string
	Next *Node
	Prev *Node
}

type QueueLinkedList struct {
	Head *Node
	Tail *Node
	Length int		
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		Head: nil,
		Tail: nil,
		Length: 0,
	}
}

func createNode(topic string, message string) *Node {
	return &Node{
		Topic: topic,
		Message: message,
		Next: nil,
	}
}

func (q *QueueLinkedList) Enqueue(topic string, message string) {
	newNode := createNode(topic, message)
	q.Length++

	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
		return
	}

	q.Tail.Next = newNode
	q.Tail = newNode
}

func (q *QueueLinkedList) PrintQueue() {
	fmt.Printf("QueueLinkedList [%d]\n", q.Length)
	for node := q.Head; node != nil; node = node.Next {
		fmt.Printf("-> [%s] %s\n", node.Topic, node.Message)
	}
}

func (q *QueueLinkedList) Dequeue() (message string, topic string) {
	if q.Length == 0 {
		return "", ""
	}

	message, topic = q.Head.Message, q.Head.Topic
	
	q.Head = q.Head.Next
	q.Tail = q.Tail.Prev

	q.Length--

	return
}

