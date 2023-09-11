package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (list *List) AddNodeToBeginning(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode

}

func (list *List) AddNodeToEnd(value int) {
	newNode := &Node{data: value}

	// need to change this after implementing addnodetobeginngin
	if list.head == nil {
		list.AddNodeToBeginning(value)
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *List) Print() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("end of list")
}
