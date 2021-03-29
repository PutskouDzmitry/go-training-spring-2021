package main

import "fmt"

type List struct {
	value int
	next *List
}

type LinkedList struct {
	head *List
}

func main(){
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	//list.Deletion()
	list.Display()
	//list.Delete(0)
	//list.Display()
	list.Search(2)
	list.Sort()
	list.Display()

}

func (L *LinkedList) Insert(key int)  {
	list := &List{
		value: key,
	}
	list.next = L.head
	L.head = list
}

func (L *LinkedList) Search(id int)  {
	cur := L.head
	var i int = 0
	for cur.next != nil && id != i {
		cur = cur.next
		i++
	}
	if cur.next == nil && id != i {
		fmt.Println("Element doesn't exist")
		return
	}
	fmt.Println(cur.value)
}

func (L *LinkedList) Deletion(){
	L.head = L.head.next
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
}

func (L *LinkedList) Sort() {
	cur := L.head
	for cur != nil{
		nextElement := cur.next
		for nextElement != nil {
			if cur.value > nextElement.value {
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