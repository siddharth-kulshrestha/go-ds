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

func bstInsertAndIsComplete() {
	lst := []int{5, 6, 4, 2, 1, 10}
	fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
	op := node.IsComplete()
	fmt.Println("IsComplete: ", op)
}

func bstInsertAndIsComplete_2() {
	lst := []int{5, 3, 7, 2, 4}
	fmt.Println("Input list: ", lst)
	node := collections.ConvSliceToBST(lst)
	op := node.IsComplete()
	fmt.Println("IsComplete: ", op)
}

func bstInsertAndIsComplete_3() {
	tcs_pos := [][]int{{5, 3, 7, 2, 4, 6}, {5, 3, 7, 2, 4, 6}, {5, 3, 7, 2, 4, 6, 8}}
	for _, tc := range tcs_pos {
		node := collections.ConvSliceToBST(tc)
		op := node.IsComplete()
		if op == true {
			fmt.Printf("tc: %v \t is complete. Test case passed!\n", tc)
		} else {
			fmt.Printf("tc: %v \t is complete. Test case failed!\n", tc)
		}
	}

	tcs_neg := [][]int{{5, 6, 4, 2, 1, 10}, {5, 4, 7, 3, 6, 8},
		{10, 5, 15, 2, 7, 12, 20, 1},
		{8, 4, 12, 2, 6, 10, 14, 1, 3, 9}, {20, 10, 30, 5, 15, 25, 35, 17}, {50, 25, 75, 10, 40, 60, 90, 5, 45}}
	for _, tc := range tcs_neg {
		node := collections.ConvSliceToBST(tc)
		op := node.IsComplete()
		if op == false {
			fmt.Printf("tc: %v \t is not complete. Test case passed!\n", tc)
		} else {
			fmt.Printf("tc: %v \t is not complete. Test case failed!\n", tc)
		}
	}

}

func init() {
	RegisterRunner(bstInsertAndThenInorder, "BST Insert -- Inorder Example")
	RegisterRunner(bstInsertAndThenPreorder, "BST Insert -- Preorder Example")
	RegisterRunner(bstInsertAndThenPostorder, "BST Insert -- Postorder Example")
	RegisterRunner(bstInsertAndHeight, "BST Insert -- Height Example")
	RegisterRunner(bstInsertAndDiameter, "BST Insert -- Diameter Example")
	RegisterRunner(bstInsertAndIsComplete, "BST Insert -- IsComplete Example")
	RegisterRunner(bstInsertAndIsComplete_2, "BST Insert -- IsComplete Example")
	RegisterRunner(bstInsertAndIsComplete_3, "BST Insert -- IsComplete Example")
}
