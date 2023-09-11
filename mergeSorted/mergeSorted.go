package mergesorted

import ()

type SortNode interface {
	GetValue() int
	GetNext() SortNode
}

type SortableList interface {
	GetHead() SortNode
	SetHead(SortNode)
	Append(int)
}

type SingleLinked struct {
	data int
	next *SingleLinked
}

type SingleList struct {
	head *SingleLinked
}

func (node *SingleLinked) GetValue() int {
	return node.data
}

func (node *SingleLinked) GetNext() SortNode {
	return node.next
}

func (list *SingleList) GetHead() SortNode {
	return list.head
}

func (list *SingleList) SetHead(node SortNode) {
	list.head = node.(*SingleLinked)
}

func (list *SingleList) Append(value int) {
	newNode := &SingleLinked{data: value, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}

type DoubleLinked struct {
	data   int
	next   *DoubleLinked
	double *DoubleLinked
}

type DoubleList struct {
	head *DoubleLinked
}

func (node *DoubleLinked) GetValue() int {
	return node.data
}

func (node *DoubleLinked) GetNext() SortNode {
	return node.next
}

func (list *DoubleList) GetHead() SortNode {
	return list.head
}

func (list *DoubleList) SetHead(node SortNode) {
	list.head = node.(*DoubleLinked)
}

func (list *DoubleList) Append(value int) {
	newNode := &DoubleLinked{data: value, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}
