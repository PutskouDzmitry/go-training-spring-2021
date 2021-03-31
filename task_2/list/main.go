package main

import (
	"fmt"
	"reflect"
)

type List struct {
	value interface{}
	next *List
}

type LinkedList struct {
	head *List
	len int
}

func main(){
	list := LinkedList{}
	list.Insert('1')
	list.Insert('2')
	list.Sort()
	list.Display()

}

func (L *LinkedList) Insert(key interface{})  {
	list := &List{
		value: key,
	}
	list.next = L.head
	L.head = list
	L.len++
}

func (L *LinkedList) Search(id int)  {
	cur := L.head
	var i int = 0
	for cur.next != nil && id != i {
		cur = cur.next
		i++
	}
	if cur.next == nil && id != i {
		fmt.Println("Element not found")
		return
	}
	fmt.Println(cur.value)
}

func (L *LinkedList) Deletion(){
	if L.head == nil {
		fmt.Println("Your list is empty")
		return
	}
	L.head = L.head.next
	L.len--
}

func (L *LinkedList) Delete(id int){
	prev := L.head
	cur := L.head
	var i int = 0
	for cur.next != nil && id != i {
		prev = cur
		cur = cur.next
		i++
	}
	if cur.next == nil && id != i {
		fmt.Println("Element doesn't exist")
		return
	}
	if cur == L.head {
		L.head = L.head.next
	} else {
		prev.next = cur.next
	}
	L.len--
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

func (L *LinkedList) Display()  {
	list := L.head
	for list != nil {
		fmt.Printf("%v ", list.value)
		list = list.next
	}
	fmt.Println()
}