package main

import (
	"github.com/siddharth-kulshrestha/go-ds/collections"
)

func main() {

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
