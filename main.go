package main

import (
	"fmt"
	"github.com/jacksongodsey/exercise2/circularlinkedlist"
	"github.com/jacksongodsey/exercise2/doublelinkedlist"
	"github.com/jacksongodsey/exercise2/linkedlist"
)

func main() {
	fmt.Println("Regular Linked List")
	list := &linkedlist.List{}
	list.AddNodeToBeginning(1)
	list.AddNodeToBeginning(2)
	list.AddNodeToEnd(3)

	list.Print()

	fmt.Println("Double Linked List")
	doublelist := &doublelinkedlist.DoubleList{}
	doublelist.AddNodeToBeginning(1)
	doublelist.AddNodeToBeginning(2)
	doublelist.AddNodeToEnd(3)

	doublelist.Print()

	fmt.Println("Circular Linked List")
	circlelist := &circularlinkedlist.CircularList{}
	circlelist.AddNodeToBeginning(4)
	circlelist.AddNodeToBeginning(1)
	circlelist.AddNodeToEnd(2)

	circlelist.Print()
}
