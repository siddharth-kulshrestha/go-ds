package main

import (
	"fmt"
	"github.com/siddharth-kulshrestha/go-ds/collections"
)

func bstInsertExample() {
	lst := []int{5, 6, 4, 2, 1, 10}
	fmt.Println("BST List: ", lst)
	node := collections.ConvSliceToBST(lst)
	fmt.Println(node)
  fmt.Println(node.Len())
  op := make([]int, 0)
  op = node.Inorder(op)
  fmt.Println(op)
  

}

func init() {
	RegisterRunner(bstInsertExample, "BST Insert Example")
}
