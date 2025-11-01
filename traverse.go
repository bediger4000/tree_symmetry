package main

// Traverse a multi-child list in preorder, shoe us the results

import (
	"fmt"
	"log"
	"os"
	"tree_symmetry/multitree"
)

func main() {
	root, err := multitree.FromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	multitree.PreorderVisitorTraverse(root, multitree.PrintVisitor)
	fmt.Println()
}
