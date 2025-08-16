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

func bstInsertAndThenPostorder() {
	lst := []int{5, 6, 4, 2, 1, 10}
	fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
	op := node.Postorder()
	fmt.Println(op)
}

func bstInsertAndThenPreorder() {
	lst := []int{5, 6, 4, 2, 1, 10}
	fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
	op := node.Preorder()
	fmt.Println(op)
}

func bstInsertAndHeight() {
	lst := []int{5, 6, 4, 2, 1, 10}
	fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
	op := node.Height()
	fmt.Println(op)
}

func bstInsertAndDiameter() {
	lst := []int{5, 6, 4, 2, 1, 10}
	fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
	op := node.Diameter()
  fmt.Println("Diameter: ", op)
}

func init() {
	RegisterRunner(bstInsertAndThenInorder, "BST Insert -- Inorder Example")
	RegisterRunner(bstInsertAndThenPreorder, "BST Insert -- Preorder Example")
	RegisterRunner(bstInsertAndThenPostorder, "BST Insert -- Postorder Example")
	RegisterRunner(bstInsertAndHeight, "BST Insert -- Height Example")
	RegisterRunner(bstInsertAndDiameter, "BST Insert -- Diameter Example")
}
