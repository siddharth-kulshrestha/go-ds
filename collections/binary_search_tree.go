package collections

import (
	"fmt"
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

func (bst *BSTNode[T]) postorder(collector []T) []T {
	if bst.Left != nil {
		collector = bst.Left.postorder(collector)
	}
	if bst.Right != nil {
		collector = bst.Right.postorder(collector)
	}
	collector = append(collector, bst.Value)
	return collector
}

func (bst *BSTNode[T]) preorder(collector []T) []T {
	collector = append(collector, bst.Value)
	if bst.Left != nil {
		collector = bst.Left.preorder(collector)
	}
	if bst.Right != nil {
		collector = bst.Right.preorder(collector)
	}
	return collector
}

func (bst *BSTNode[T]) Preorder() []T {
	collector := make([]T, 0)
	collector = bst.preorder(collector)
	return collector
}

func (bst *BSTNode[T]) Postorder() []T {
	collector := make([]T, 0)
	collector = bst.postorder(collector)
	return collector
}

func (bst *BSTNode[T]) diameter(ht int) (int, int) {
	var dl, dr, hr, hl int
	if bst.Left == nil && bst.Right == nil {
		return 0, -1
	}
	if bst.Left != nil {
		dl, hl = bst.Left.diameter(ht)
	}
	if bst.Right != nil {
		dr, hr = bst.Right.diameter(ht)
	}

	dmax := max(dl, dr, hr+hl+2)
	hmax := max(hl, hr)
	return dmax, hmax + 1
}

func (bst *BSTNode[T]) Diameter() int {
	dm, ht := bst.diameter(0)
	fmt.Println("Height of tree in diameter computation: ", ht)
	return dm
}

func (bst *BSTNode[T]) height(ht int) int {
	var hl, hr int
	if bst.Left == nil && bst.Right == nil {
		return 0
	}
	if bst.Left != nil {
		hl = bst.Left.height(ht)
	}
	if bst.Right != nil {
		hr = bst.Right.height(ht)
	}
	max := max(hl, hr)
	return max + 1
}

func (bst *BSTNode[T]) Height() int {
	return bst.height(0)
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
