package main

import "fmt"

type Element struct {
	value int
	next *Element
}

type Queue struct {
	head *Element
	tail *Element
	size int
	maxSize int
}

func main() {
	queue := Queue{
		maxSize: 10,
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.IsEmpty()
	queue.Peek()
	queue.Dequeue()
	queue.Dequeue()
	queue.Peek()
}

func (Q *Queue) IsFull()  {
	if Q.maxSize == Q.size {
		fmt.Println("Queue is full")
	} else {
		fmt.Println("Queue isn't full")
	}
}

func (Q *Queue) IsEmpty()  {
	if Q.size == 0 {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Queue isn't empty")
	}
}

func (Q *Queue) Peek()  {
	if Q.size != 0 {
		fmt.Println(Q.head.value)
	} else {
		fmt.Println("Queue is empty")
	}
}

func (Q *Queue) Enqueue(value int)  {
	if Q.size == Q.maxSize {
		fmt.Println("Queue is full")
		return
	}
	element := &Element{
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

func (Q *Queue) Dequeue()  {
	element := Element{}
	if Q.head != nil {
		element.value = Q.head.value
		if Q.head.next != nil {
			Q.head = Q.head.next
		} else {
			Q.head = nil
			Q.tail = nil
		}
		Q.size--
	}
}


