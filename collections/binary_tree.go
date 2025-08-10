package collections

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T any](value T) *TreeNode[T] {
	return &TreeNode[T]{
		Value: value,
	}
}
