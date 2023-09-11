package circularlinkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type CircularList struct {
	head *Node
	tail *Node
}

func (list *CircularList) AddNodeToBeginning(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		newNode.next = newNode // Make it point to itself
		newNode.prev = newNode // Make it point to itself
		return
	}

	newNode.next = list.head
	list.head.prev = newNode

	list.tail.next = newNode
	newNode.prev = list.tail

	list.head = newNode
}

func (list *CircularList) AddNodeToEnd(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.AddNodeToBeginning(value)
		return
	}

	current := list.head
	for current.next != list.head {
		current = current.next
	}

	current.next = newNode
	newNode.next = list.head
	list.head.prev = newNode
	newNode.prev = current

	list.tail = newNode
}

func (list *CircularList) Print() {
	current := list.head.prev
	beginning := list.head
	fmt.Print("end of list -> ")
	for current != beginning {
		fmt.Print(current.data, " -> ")
		current = current.prev
	}
	fmt.Print(current.data, " -> ")
	fmt.Println("start of list")
}
