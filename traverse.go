package main

// Traverse a multi-child list in preorder, shoe us the results

import (
	"fmt"
	"os"
	"tree_symmetry/multitree"
)

func main() {
	root := multitree.FromString(os.Args[1])
	multitree.PreorderVisitorTraverse(root, multitree.PrintVisitor)
	fmt.Println()
}
