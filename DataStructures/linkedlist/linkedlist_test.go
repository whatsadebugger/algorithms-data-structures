package linkedlist_test

import (
	"testing"

	"github.com/whatsadebugger/algorithms-data-structures/DataStructures/linkedlist"
)

func TestLinkedList(t *testing.T) {
	head := linkedlist.NewList()
	head.Append(10)
	head.Append(10)
	head = head.Push(1)
	linkedlist.Print(head)
}

func TestInsert(t *testing.T) {

}
func TestUpdate(t *testing.T) {

}
func TestDelete(t *testing.T) {

}
