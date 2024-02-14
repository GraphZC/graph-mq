package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/GraphZC/graph-mq/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmptyLinkedList(t *testing.T) {
	newList := model.NewQueueLinkedList()

	mockList := &model.QueueLinkedList{
		Head: nil,
		Tail: nil,
		Length: 0,
	}

	assert.True(t, reflect.DeepEqual(newList, mockList)) 
}

func TestLengthWhenListIsEmpty(t *testing.T) {
	list := model.NewQueueLinkedList()

	assert.Equal(t, 0, list.Length)
}

func TestLengthEnqueueWhenListIsEmpty(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "msg")

	assert.Equal(t, 1, list.Length)
}

func TestEnqueueWhenListIsNotEmpty(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "test-1")
	list.Enqueue("topic", "test-2")

	assert.Equal(t, 2, list.Length)
}

func TestDequeueWhenListIsEmpty(t *testing.T) {
	list := model.NewQueueLinkedList()

	msg, topic := list.Dequeue()

	assert.True(t, msg == "" && topic == "")
}

func TestLenghtDequeueWhenListHasOneQueue(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "test")
	list.Dequeue()
	
	assert.Equal(t, 0, list.Length)
}

func TestMessageDequeueWhenListHasOneQueue(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "test")
	msg, _ := list.Dequeue()
	
	assert.Equal(t, "test", msg)
}

func TestTopicDequeueWhenListHasOneQueue(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "test-1")
	list.Enqueue("topic", "test-2")

	msg, topic := list.Dequeue()
	fmt.Println(msg, topic)
	
	assert.Equal(t, "topic", topic)
}

func TestDequeueLenghtWhenListHasMoreOneQueue(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "test-1")
	list.Enqueue("topic", "test-2")

	list.Dequeue()
	
	assert.Equal(t, 1, list.Length)
}

func TestDequeueMessageWhenListHasMoreOneQueue(t *testing.T) {
	list := model.NewQueueLinkedList()

	list.Enqueue("topic", "test-1")
	list.Enqueue("topic", "test-2")

	msg, _ := list.Dequeue()
	
	assert.Equal(t, "test-1", msg)
}
