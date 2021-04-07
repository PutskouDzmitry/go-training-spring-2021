package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_GetPosition(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList(1,2,3)
	node := Node{
		value: 2,
	}
	element, _ := list.GetPosition(1)
	assert.Equal(node.value, element.value, fmt.Sprintf("%v and %v not equal, but expected:", node.value, element.value))
}

func TestLinkedList_GetPosition2(t *testing.T) {
	assert := assert.New(t)
	var index = 6
	list := NewLinkedList(1,2,3)
	_, err := list.GetPosition(index)
	assert.EqualError(err, fmt.Sprintf("your id is greater than length of LinkedList id: %v",index))
}

func TestNewLinkedList(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList(1,2,3)
	expectedList := LinkedList{}
	expectedList.Insert(1)
	expectedList.Insert(2)
	expectedList.Insert(3)
	for i := 0;  i < list.GetLen(); i++ {
		expected, _ := expectedList.Search(i)
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
	}
}

func TestLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList(1,2,3,4,7,5)
	expectedArr := []int{1,2,3,4,5,7}
	list.Sort()
	for i := 0;  i < list.GetLen(); i++ {
		expected := expectedArr[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
	}
}

func TestLinkedList_GetLen(t *testing.T) {
	list := NewLinkedList(1,2,3)
	assert := assert.New(t)
	expected := 3
	actual := list.len
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestLinkedList_Search1(t *testing.T) {
	list := NewLinkedList("a","b","c","d")
	assert := assert.New(t)
	list.Insert("f")
	expectedList := []string {"f","d","c","b","a"}
	for i := 0;  i < list.GetLen(); i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestLinkedList_Search2(t *testing.T) {
	list := NewLinkedList("a","b","c","d")
	assert := assert.New(t)
	_, err := list.Search(9)
	assert.EqualError(err, "your id is greater than length of LinkedList id: 9")
}

func TestLinkedList_Deletion1(t *testing.T) {
	list := NewLinkedList("a","b","c","d")
	assert := assert.New(t)
	list.Deletion()
	list.Deletion()
	expectedList := []string {"b","a"}
	for i := 0;  i < list.GetLen(); i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestLinkedList_Deletion2(t *testing.T) {
	list := NewLinkedList("a","b","c","d")
	assert := assert.New(t)
	list.Deletion()
	list.Deletion()
	list.Deletion()
	list.Deletion()
	list.Deletion()
	err := list.Deletion()
	assert.EqualError(err, "the list is empty")
}

func TestLinkedList_Delete1(t *testing.T) {
	list := NewLinkedList("a","b","c","d")
	assert := assert.New(t)
	list.Delete(1)
	list.Delete(list.GetLen() - 1)
	expectedList := []string {"d","b"}
	for i := 0; i < list.GetLen(); i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
	}
}

func TestLinkedList_Delete2(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < list.GetLen() + 85; i++ {
		list.Delete(0)
	}
	err := list.Delete(0)
	assert.EqualError(err, "your list is empty")
}

