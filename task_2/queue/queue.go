package queue

import (
	"errors"
	"fmt"
	"reflect"
)

// Node represents node of data struct queue.
type Node  struct {
	value interface{} // an value that is in the node
	next *Node // pointer to the next node
}

// Queue represents the implementation of queue.
type Queue struct {
	head *Node // the head of the queue
	tail *Node // the tail of the queue
	size int  // the size of the queue
	capacity int // the max capacity of the queue
}

func (Q *Queue) GetSize () int {
	return Q.size
}

func NewQueue(max int) *Queue {
	return &Queue{capacity : max}
}

// typeDefinition compare two elements of Queue
// If i1 > i2, returns true
// If i1 < i2, returns false
// If i1 and i2 don't equals in type, then an error message will be printed and typeDefinition will be closed
// If type of i1 and type of i2 don't equals: int, float64, string, rune, uint and byte, then an error message will be printed and typeDefinition will be closed
func typeDefinition(i1 interface{}, i2 interface{}) (bool, error) {
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int), nil
	case float64:
		return i1.(int) > i2.(int), nil
	case string:
		return i1.(int) > i2.(int), nil
	case rune:
		return i1.(int) > i2.(int), nil
	case uint:
		return i1.(int) > i2.(int), nil
	case byte:
		return i1.(int) > i2.(int), nil
	default:
		return false, fmt.Errorf("this type isn't supported: %v", reflect.TypeOf(i1))
	}
}

// Sort sort Queue using bubble sort algorithm
func (Q *Queue) Sort() error {
	cur := Q.head
	for cur != nil{
		nextElement := cur.next
		for nextElement != nil {
			if t, err := typeDefinition(cur.value, nextElement.value); t {
				if err != nil {
					return err
				}
				temp := cur.value
				cur.value = nextElement.value
				nextElement.value = temp
			}
			nextElement = nextElement.next
		}
		cur = cur.next
	}
	return nil
}

// IsFull checks if the queue is full
func (Q *Queue) IsFull() bool {
	return Q.capacity <= Q.size
}

// IsEmpty checks if the queue is empty
func (Q *Queue) IsEmpty() bool {
	return Q.size == 0
}
// Peek gets the value of the front of the queue without removing it
// if queue is empty, then an error message will be printed and Peek will be closed
func (Q *Queue) Peek() (interface{}, error) {
	if Q.IsEmpty() {
		return nil, errors.New("queue is empty. You can't display element :(")
	}
	return Q.head.value, nil
}

// Enqueue adds an element to the end of the queue
// if queue is full, then an error message will be printed and Enqueue will be closed
func (Q *Queue) Enqueue(value interface{}) error {
	if Q.IsFull() {
		return errors.New("queue is full. You can't add element in your queue(")
	}
	element := &Node {
		value: value,
	}
	if Q.head == nil {
		Q.tail = element
		Q.head = element
	} else {
		Q.tail.next = element
		Q.tail = element
	}
	Q.size++
	return nil
}

// Dequeue is removes an element from the front of the queue
// if element in Queue equals nil, then an error message will be printed and Dequeue will be closed
func (Q *Queue) Dequeue() error {
	if Q.head == nil {
		return errors.New("your queue is empty and impossible delete element")
	}
	element := Node {}
	element.value = Q.head.value
	if Q.head.next != nil {
		Q.head = Q.head.next
	} else {
		Q.head = nil
		Q.tail = nil
	}
	Q.size--
	return nil
}


