package main

import (
	"github.com/siddharth-kulshrestha/go-ds/collections"
)

func linkListExamples() {

	sl1 := []int{2, 3, 4, 5}
	sl2 := []int{6, 5, 1, 2, 4, 5}
	sl3 := []int{9, 21, 23, 123, 53, 2}
	sl4 := []int{9}
	sl5 := []int{10}
	sl6 := []any{"str1", 1, 2}

	ll1 := collections.ConvSliceToLinkList(sl1)
	ll2 := collections.ConvSliceToLinkList(sl2)
	ll3 := collections.ConvSliceToLinkList(sl3)
	ll4 := collections.ConvSliceToLinkList(sl4)
	ll5 := collections.ConvSliceToLinkList(sl5)
	ll6 := collections.ConvSliceToLinkList(sl6)

	ll1.PrintTillEnd()
	ll2.PrintTillEnd()
	ll3.PrintTillEnd()
	ll4.PrintTillEnd()
	ll5.PrintTillEnd()
	ll6.PrintTillEnd()

}

func deleteNthNodeFromEnd[T any](head *collections.LinkedListNode[T], n int) *collections.LinkedListNode[T] {
	// Using two pointer approach
	// we will move the end pointer by n steps, then start pointer will start moving thus creating a window of start --> end pointer.
	// and as soon as end pointer reaches the end of the list we will remove the start pointer's next element.
	if n == 0 {
		return head
	}

	var start, end *collections.LinkedListNode[T]

	start = head
	end = head
	for _ = range n {
		if end.Next == nil {
			continue
		}
		end = end.Next
	}

	// Now end is n steps ahead

	if end.Next == nil {
		// which means either there is only a single element or n is equal to length of LinkList, which means delete the 0th node.
		head = start.Next
		return head
	}

	for end.Next != nil {
		end = end.Next
		start = start.Next
	}

	// end has reached the last.N	// end has reached the last.Next.
	// if start == end {
	// 	// this means that there is either only single node & n is 1
	// 	head = nil

	// }

	if start.Next != nil {
		start.Next = start.Next.Next
	}
	return head
}

func deleteNthNodeFromEnd_op[T any](head *collections.LinkedListNode[T], n int) {
	head = deleteNthNodeFromEnd(head, n)
	head.PrintTillEnd()
}

func deleteNthNodeFromEndLinkList_Example() {
	sl1 := []int{2, 3, 4, 5}
	sl2 := []int{6, 5, 1, 2, 4, 5}
	sl3 := []int{9, 21, 23, 123, 53, 2}
	sl4 := []int{9}
	sl5 := []int{10}
	sl6 := []any{"str1", 1, 2}

	ll1 := collections.ConvSliceToLinkList(sl1)
	ll1.PrintTillEnd()
	deleteNthNodeFromEnd_op(ll1, 2)
	ll2 := collections.ConvSliceToLinkList(sl2)
	ll2.PrintTillEnd()
	deleteNthNodeFromEnd_op(ll2, 2)
	ll3 := collections.ConvSliceToLinkList(sl3)
	ll3.PrintTillEnd()
	deleteNthNodeFromEnd_op(ll3, 2)
	ll4 := collections.ConvSliceToLinkList(sl4)
	ll4.PrintTillEnd()
	deleteNthNodeFromEnd_op(ll4, 1)
	ll5 := collections.ConvSliceToLinkList(sl5)
	ll5.PrintTillEnd()
	deleteNthNodeFromEnd_op(ll5, 1)
	ll6 := collections.ConvSliceToLinkList(sl6)
	ll6.PrintTillEnd()
	deleteNthNodeFromEnd_op(ll6, 3)

}

func init() {
	RegisterRunner(linkListExamples, "LinkList Basic Examples")
	RegisterRunner(deleteNthNodeFromEndLinkList_Example, "Delete Nth Node from end using two pointer approach")
}
