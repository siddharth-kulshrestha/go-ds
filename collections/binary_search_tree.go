package collections

type BSTNode[T comparable] struct {
	*TreeNode[T]
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode[T comparable](value T) *BSTNode[T] {
	return &BSTNode[T]{
		TreeNode: NewTreeNode(value),
	}
}

func (bst *BSTNode[T]) Insert(value T) {
	if value < bst.Value && bst.Left != nil {
		if bst.Left != nil {
			bst.Left.Insert(value)
		} else {
			bst.Left = NewBST
		}
	} else if value >= bst.Value && bst.Right != nil {
		bst.Right.Insert(value)
	}
}
