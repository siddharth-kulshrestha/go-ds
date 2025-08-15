package collections

import (
	"golang.org/x/exp/constraints"
)

type BSTNode[T constraints.Ordered] struct {
	*TreeNode[T]
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

func NewBSTNode[T constraints.Ordered](value T) *BSTNode[T] {
	return &BSTNode[T]{
		TreeNode: NewTreeNode(value),
	}
}

func (bst *BSTNode[T]) Insert(value T) {
	if value < bst.Value {
		if bst.Left != nil {
			bst.Left.Insert(value)
		} else {
			bst.Left = NewBSTNode(value)
		}
	} else if value >= bst.Value {
		if bst.Right != nil {
			bst.Right.Insert(value)
		} else {
			bst.Right = NewBSTNode(value)
		}
	}
}

func (bst *BSTNode[T]) len(start int) int {
	var ll, lr int
	if bst.Left != nil {
		ll = bst.Left.len(start)
	}
	if bst.Right != nil {
		lr = bst.Right.len(start)
	}
	if bst.Left == nil && bst.Right == nil {
		return start + 1
	}
	return ll + lr + 1
}

func (bst *BSTNode[T]) Len() int {
	return bst.len(0)
}

func (bst *BSTNode[T]) inorder(collector []T) []T {
	if bst.Left != nil {
		collector = bst.Left.inorder(collector)
	}
	collector = append(collector, bst.Value)
	if bst.Right != nil {
		collector = bst.Right.inorder(collector)
	}
	return collector
}

func (bst *BSTNode[T]) Inorder() []T {
  collector := make([]T, 0)
  collector = bst.inorder(collector)
  return collector
}

func ConvSliceToBST[T constraints.Ordered](lst []T) *BSTNode[T] {
	if len(lst) == 0 {
		return nil
	}
	node := NewBSTNode(lst[0])
	for _, v := range lst[1:] {
		node.Insert(v)
	}
	return node
}

// func (bst *BSTNode[T]) PrintInordered() string {
//

// }
