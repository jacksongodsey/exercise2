package doublelinkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoubleList struct {
	head *Node
}

func (list *DoubleList) AddNodeToBeginning(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
		return
	}

	newNode.next = list.head
	list.head.prev = newNode
	list.head = newNode

}

func (list *DoubleList) AddNodeToEnd(value int) {
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
	newNode.prev = current
}

func (list *DoubleList) Print() {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.prev
	}
	fmt.Println("start of list")
}
