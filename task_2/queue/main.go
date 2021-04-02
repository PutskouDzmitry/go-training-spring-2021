package main

import (
	"fmt"
	"os"
	"reflect"
)

// Node represents node of data struct queue.
type Node  struct {
	value interface{}
	next *Node
}

// Queue represents the implementation of queue.
type Queue struct {
	head *Node
	tail *Node
	size int
	maxSize int
}

func main() {
	queue := Queue{
		maxSize: 2,
	}
	queue.IsEmpty()
	queue.IsFull()
	queue.Enqueue(412)
	queue.Dequeue()
}

// typeDefinition compare two elements of Queue
// If i1 > i2, returns true
// If i1 < i2, returns false
// If i1 and i2 don't equals in type, then an error message will be printed and typeDefinition will be closed
// If type of i1 and type of i2 don't equals: int, float64, string, rune, uint and byte, then an error message will be printed and typeDefinition will be closed
func typeDefinition(i1 interface{}, i2 interface{}) bool {
	if reflect.TypeOf(i1) != reflect.TypeOf(i2) {
		fmt.Println("Type incompatibility (")
		os.Exit(1)
	}
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int)
	case float64:
		return i1.(float64) > i2.(float64)
	case string:
		return i1.(string) > i2.(string)
	case rune:
		return i1.(rune) > i2.(rune)
	case uint:
		return i1.(uint) > i2.(uint)
	case byte:
		return i1.(byte) > i2.(byte)
	default:
		fmt.Println("This type isn't supported")
		os.Exit(1)
	}
	return false
}

// Sort sort Queue using bubble sort algorithm
func (Q *Queue) Sort() {
	cur := Q.head
	for cur != nil{
		nextElement := cur.next
		for nextElement != nil {
			if typeDefinition(cur.value, nextElement.value) {
				temp := cur.value
				cur.value = nextElement.value
				nextElement.value = temp
			}
			nextElement = nextElement.next
		}
		cur = cur.next
	}
}

// IsFull checks if the queue is full
func (Q *Queue) IsFull() bool {
	return Q.maxSize <= Q.size
}

// IsEmpty checks if the queue is empty
func (Q *Queue) IsEmpty() bool {
	return Q.size == 0
}
// Peek gets the value of the front of the queue without removing it
// if queue is empty, then an error message will be printed and Peek will be closed
func (Q *Queue) Peek() interface{} {
	if Q.IsEmpty() {
		fmt.Println("Queue is empty. You can't display element :(")
		os.Exit(1)
	}
	return Q.head.value
}

// Enqueue adds an element to the end of the queue
// if queue is full, then an error message will be printed and Enqueue will be closed
func (Q *Queue) Enqueue(value interface{})  {
	if Q.IsFull() {
		fmt.Println("Queue is full. You can't add element in your queue(")
		os.Exit(1)
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
}

// Dequeue is removes an element from the front of the queue
// if element in Queue equals nil, then an error message will be printed and Dequeue will be closed
func (Q *Queue) Dequeue()  {
	if Q.head == nil {
		fmt.Println("Your queue is empty and impossible delete element")
		os.Exit(1)
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
}


