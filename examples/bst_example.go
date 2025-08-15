package main

import (
	"fmt"
	"github.com/siddharth-kulshrestha/go-ds/collections"
)

func bstInsertAndThenInorder() {
	lst := []int{5, 6, 4, 2, 1, 10}
  fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
  op := node.Inorder()
  fmt.Println(op)
}



func init() {
	RegisterRunner(bstInsertAndThenInorder, "BST Insert Example")
}
