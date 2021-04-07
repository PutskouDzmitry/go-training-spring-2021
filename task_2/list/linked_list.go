package linked_list

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Node represents node of data struct linked list.
type Node struct {
	value interface{} // an value that is in the node
	next *Node // pointer to the next node
}

// LinkedList represents the implementation of singly linked list.
type LinkedList struct {
	head *Node // the head of the list
	len int // the Len of the list
}


// GetLen get len of LinkedList
func (L LinkedList) GetLen() int {
	return L.len
}

// NewLinkedList create an empty LinkedList
func NewLinkedList(values ...interface{}) *LinkedList {
	list := &LinkedList{}
	for _, value := range values {
		list.Insert(value)
	}
	return list
}

// increaseLen increases the length of the LinkedList
func (L *LinkedList) increaseLen() {
	L.len++
}

// decreaseLen decreases the length of the LinkedList
func (L *LinkedList) decreaseLen() {
	L.len--
}

// Insert Adds an element at the beginning of the list.
func (L *LinkedList) Insert(key interface{})  {
	list := &Node{
		value: key,
	}
	list.next = L.head
	L.head = list
	L.increaseLen()
}

// Search Searches an element using the id.
// If Search doesn't find an element in LinkedList or id >= length of LinkedList, then then an error message will be printed and Search will be closed
func (L *LinkedList) Search(id int) (interface{}, error) {
	cur, errorCur := L.GetPosition(id)
	if errorCur != nil {
		return nil, errorCur
	}
	if cur == nil {
		return nil, errors.New("element not found")
	}
	return cur.value, nil
}

// Deletion deletes an element at the beginning of the list.
// If Length of LinkedList equals 0, then then an error message will be printed and Deletion will be closed
func (L *LinkedList) Deletion() error{
	if L.len == 0 {
		return errors.New("the list is empty")
	}
	L.head = L.head.next
	L.decreaseLen()
	return nil
}

// Delete deletes an element using the id.
// If Length of LinkedList equals 0 or if Delete, then then an error message will be printed and Delete will be closed
func (L *LinkedList) Delete(id int) error{
	if L.len == 0 {
		return errors.New("your list is empty")
	}
	if id == 0 {
		return L.Deletion()
	} else {
		cur, errorCur := L.GetPosition(id)
		if errorCur != nil {
			return errorCur
		}
		prev, errorPrev := L.GetPosition(id - 1)
		if errorPrev != nil {
			return errorPrev
		}
		if cur == L.head {
			L.head = L.head.next
		} else {
			prev.next = cur.next
		}
		L.decreaseLen()
	}
	return nil
}

// getPosition gets a node by the id.
func (L *LinkedList) GetPosition(id int) (*Node, error) {
	if id < 0 {
		return nil, fmt.Errorf("your id is incorrect id: %v", id)
	}
	if id >= L.len {
		return nil, fmt.Errorf("your id is greater than length of LinkedList id: %v", id)
	}
	element := L.head
	for i := 0; i < id; i++ {
		element = element.next
	}
	return element, nil
}

// typeDefinition compare two elements of LinkedList
// If i1 > i2, returns true
// If i1 < i2, returns false
// If i1 and i2 don't equals in type, then an error message will be printed and typeDefinition will be closed
// If type of i1 and type of i2 don't equals: int, float64, string, rune, uint and byte, then an error message will be printed and typeDefinition will be closed
func typeDefinition(i1 interface{}, i2 interface{}) bool {
	if reflect.TypeOf(i1) != reflect.TypeOf(i2) {
		fmt.Println("type incompatibility (")
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

// Sort sort LinkedList using bubble sort algorithm
func (L *LinkedList) Sort() {
	cur := L.head
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

// Display displays the complete list to the console.
// If length of LinkedList equals 0, then then an error message will be printed and Display will be closed
func (L *LinkedList) Display() error {
	if L.len == 0 {
		return errors.New("data output isn't possible because list is empty")
	}
	list := L.head
	for list != nil {
		fmt.Printf("%v ", list.value)
		list = list.next
	}
	fmt.Println()
	return nil
}