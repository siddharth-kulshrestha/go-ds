package collections

import (
	"fmt"
)

type LinkedListNode[T any] struct {
	Next  *LinkedListNode[T]
	Value T
}

func NewLinkedListNode[T any](value T) *LinkedListNode[T] {
	return &LinkedListNode[T]{
		Next:  nil,
		Value: value,
	}
}

func (lln *LinkedListNode[T]) Add(value T) {
	nln := NewLinkedListNode(value)
	lln.Next = nln
}

func (lln *LinkedListNode[T]) PrintTillEnd() {
	tmp := lln
	for tmp != nil {
		fmt.Print(tmp.Value, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}

func ConvSliceToLinkList[T any](lst []T) *LinkedListNode[T] {
	var head *LinkedListNode[T]
	var prev *LinkedListNode[T]
	for _, v := range lst {
		tmp := NewLinkedListNode(v)
		if head == nil {
			head = tmp
			prev = tmp
			continue
		}
		prev.Next = tmp
		prev = tmp
	}
	return head
}
