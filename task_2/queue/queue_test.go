package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Dequeue()
	expected := []int {2,3,4}
	for i := 0;  i < queue.size; i++ {
		value := queue.head
		for j := 0; j < i; j++ {
			value = value.next
		}
		actual := value.value
		assert.Equal(expected[i], actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
	}
}

func TestQueue_Dequeue2(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(3)
	err := queue.Dequeue()
	assert.EqualError(err, "your queue is empty and impossible delete element")
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	expected := []int{1,2,3,4}
	for i := 0;  i < queue.size; i++ {
		value := queue.head
		for j := 0; j < i; j++ {
			value = value.next
		}
		actual := value.value
		assert.Equal(expected[i], actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))

	}
}

func TestQueue_Enqueue2(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	err := queue.Enqueue(4)
	assert.EqualError(err, "queue is full. You can't add element in your queue(")
}

func TestQueue_GetSize(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	assert.Equal(0, queue.GetSize(), fmt.Sprintf("%v and %v not equal, but expected:", 0, queue.GetSize()))
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	assert.Equal(true, queue.IsEmpty(), fmt.Sprintf("%v and %v not equal, but expected:", true, queue.IsEmpty()))
}

func TestQueue_IsEmpty2(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	queue.Enqueue(1)
	assert.Equal(false, queue.IsEmpty(), fmt.Sprintf("%v and %v not equal, but expected:", false, queue.IsEmpty()))
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	assert.Equal(false, queue.IsFull(), fmt.Sprintf("%v and %v not equal, but expected:", false, queue.IsFull()))
}

func TestQueue_IsFull2(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(1)
	queue.Enqueue(1)
	assert.Equal(true, queue.IsFull(), fmt.Sprintf("%v and %v not equal, but expected:", true, queue.IsFull()))
}

func TestQueue_Sort(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Sort()
	expected := []int{1,2,3,4}
	for i := 0;  i < queue.size; i++ {
		value := queue.head
		for j := 0; j < i; j++ {
			value = value.next
		}
		actual := value.value
		assert.Equal(expected[i], actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))

	}
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	actual, _ := queue.Peek()
	assert.Equal(4, actual, fmt.Sprintf("%v and %v not equal, but expected:", 4, actual))
}

func TestQueue_Peek2(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	_, err := queue.Peek()
	assert.EqualError(err, "queue is empty. You can't display element :(")
}


func TestNewQueue(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(10)
	assert.Equal(10, queue.capacity, fmt.Sprintf("%v and %v not equal, but expected:", 10, queue.capacity))
}