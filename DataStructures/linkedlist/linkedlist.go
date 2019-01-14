package linkedlist

import "fmt"

// Node is a data structure which points to the next node in a linked list
type Node struct {
	data int
	Next *Node
}

// List is a linked list
type List struct{}

// NewList returns a new list
func NewList() *Node {
	return &Node{}
}

// Push to beginning of list
func (head *Node) Push(d int) *Node {
	return NewNode(d, head)
}

// Append to end of the list
func (head *Node) Append(d int) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = NewNode(d, nil)
}

// Update updates item at given position
func Update() {}

// Delete deletes an item at given position
func Delete() {}

// NewNode create node
func NewNode(d int, next *Node) *Node {
	return &Node{data: d, Next: next}
}

// Print the list at the given node as head
func Print(current *Node) {
	for current != nil {
		fmt.Println("Node: ", current.data)
		current = current.Next
	}
}
