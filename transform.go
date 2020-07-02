package main

// Perform Knuth transform on a multitree

import (
	"fmt"
	"os"
	"tree_symmetry/multitree"
	"tree_symmetry/tree"
)

func main() {
	if os.Args[1] == "-r" {
		root := multitree.FromString(os.Args[2])
		binaryTree := knuthTransform(root)
		copyTree := knuthUnTransform(binaryTree)

		fmt.Printf("digraph g1 {\n")
		fmt.Printf("subgraph cluster_0 {\n\tlabel=\"original\"\n")
		multitree.DrawPrefixed(os.Stdout, root, "original")
		fmt.Printf("\n}\n")
		fmt.Printf("subgraph cluster_1 {\n\tlabel=\"copy\"\n")
		multitree.DrawPrefixed(os.Stdout, copyTree, "copy")
		fmt.Printf("\n}\n")
		fmt.Printf("\n}\n")

		return
	}

	root := multitree.FromString(os.Args[1])
	fmt.Printf("digraph g1 {\n")
	fmt.Printf("subgraph cluster_0 {\n\tlabel=\"multitree\"\n")
	multitree.DrawPrefixed(os.Stdout, root, "a")
	fmt.Printf("\n}\n")
	binaryTree := knuthTransform(root)
	fmt.Printf("subgraph cluster_1 {\n\tlabel=\"binary tree\"\n")
	tree.DrawPrefixed(os.Stdout, binaryTree, "b")
	fmt.Printf("\n}\n")
	fmt.Printf("\n}\n")
}

// knuthTransform turns a k-ary multitree into a binary tree
// with a special format.
func knuthTransform(root *multitree.Node) *tree.NumericNode {
	newNode := &tree.NumericNode{Data: root.Data}
	if l := len(root.Children); l > 0 {
		newNode.Left = knuthTransform(root.Children[0])
		for i := range root.Children[1:] {
			newChild := knuthTransform(root.Children[l-1-i])
			newChild.Right = newNode.Left.Right
			newNode.Left.Right = newChild
		}
	}

	return newNode
}

// knuthUnTransform turns a binary tree into a  k-ary multitree
func knuthUnTransform(node *tree.NumericNode) *multitree.Node {
	newNode := &multitree.Node{Data: node.Data}

	if node.Left == nil {
		return newNode
	}

	for head := node.Left; head != nil; head = head.Right {
		child := knuthUnTransform(head)
		newNode.Children = append(newNode.Children, child)
	}

	return newNode
}
