package main

// Perform Knuth transform on a multitree

import (
	"fmt"
	"os"
	"tree_symmetry/multitree"
	"tree_symmetry/tree"
)

func main() {
	root := multitree.FromString(os.Args[1])
	fmt.Printf("digraph g1 {\n")
	fmt.Printf("subgraph cluster_0 {\n\tlabel=\"multitree\"\n")
	multitree.DrawPrefixed(os.Stdout, root, "a")
	fmt.Printf("\n}\n")
	binaryTree := transform(root)
	fmt.Printf("subgraph cluster_1 {\n\tlabel=\"binary tree\"\n")
	tree.DrawPrefixed(os.Stdout, binaryTree, "b")
	fmt.Printf("\n}\n")
	fmt.Printf("\n}\n")
}

// knuthTransform turns a k-ary multitree into a binary tree
// with a special format.
func knuthTransform(root *multitree.Node) *tree.NumericNode {
	broot := &tree.NumericNode{Data: root.Data}
	l := len(root.Children)
	if l > 0 {
		broot.Left = transform(root.Children[0])
		var list *tree.NumericNode
		for i := range root.Children[1:] {
			child := knuthTransform(root.Children[l-1-i])
			child.Right = list
			list = child
		}
		broot.Left.Right = list
	}

	return broot
}
