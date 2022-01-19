package singly

import "fmt"

// Node is a data structure which points to the next node in a linked list
type Node struct {
	data int
	Next *Node
}

// List is a linked list
type List struct {
	head *Node
}

// NewList returns a new list
func NewList() *List {
	return &List{}
}

// Push to beginning of list
func (list *List) Push(d int) {
	list.head = NewNode(d, list.head)
}

// Append to end of the list
func (list *List) Append(d int) {
	n := NewNode(d, nil)

	if list.head == nil {
		list.head = n
		return
	}

	cur := list.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
}

// Search returns the index of the value in the list and -1 if not found
func (list *List) Search(val int) int {
	cur := list.head
	for i := 0; cur != nil; i, cur = i+1, cur.Next {
		if cur.data == val {
			return i
		}
	}
	return -1
}

// NewNode create node
func NewNode(d int, next *Node) *Node {
	return &Node{data: d, Next: next}
}

// Print the list at the given node as head
func (list *List) Print() {
	cur := list.head
	for cur != nil {
		fmt.Println("Node: ", cur.data)
		cur = cur.Next
	}
}
