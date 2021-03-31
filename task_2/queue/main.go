package main

import (
	"fmt"
	"reflect"
)

type Element struct {
	value interface{}
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
		maxSize: 2,
	}
	queue.IsEmpty()
	queue.IsFull()
	queue.Enqueue(412)
	queue.Enqueue(321)
	queue.Sort()
	queue.Peek()
}

func typeDefinition(i1 interface{}, i2 interface{}) bool {
	if reflect.TypeOf(i1) != reflect.TypeOf(i2) {
		panic("Type incompatibility (")
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
		panic("This type isn't supported")
	}
}

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
		fmt.Println("Queue is empty. You can't display element :(")
	}
}

func (Q *Queue) Enqueue(value interface{})  {
	if Q.size == Q.maxSize {
		fmt.Println("Queue is full. You can't add element in your queue(")
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
	} else {
		fmt.Println("Error. Your queue is empty and you can't delete element")
	}
}


