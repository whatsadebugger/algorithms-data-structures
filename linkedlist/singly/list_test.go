package singly_test

import (
	"testing"

	"github.com/whatsadebugger/algorithms-data-structures/linkedlist/singly"
)

func TestLinkedList(t *testing.T) {
	l := singly.NewList()
	l.Append(10)
	l.Append(12)
	l.Push(1)
	l.Push(2)
	l.Append(11)
	l.Print()
	i := l.Search(11)
	if i != 4 {
		t.Errorf("want 4 but got %d", i)
	}
}
