package main

import (
	"fmt"
	"os"
	"tree_symmetry/multitree"
)

func main() {
	root := multitree.FromString(os.Args[1])

	fmt.Printf("digraph g1 {\n")
	fmt.Printf("subgraph cluster_0 {\n\tlabel=\"before\"\n")
	multitree.DrawPrefixed(os.Stdout, root, "a")
	fmt.Printf("\n}\n")
	rotated := invert(root)
	fmt.Printf("subgraph cluster_1 {\n\tlabel=\"after\"\n")
	multitree.DrawPrefixed(os.Stdout, rotated, "b")
	fmt.Printf("\n}\n")
	fmt.Printf("\n}\n")
}

func invert(node *multitree.Node) *multitree.Node {
	if node == nil {
		return nil
	}
	newnode := &multitree.Node{Data: node.Data}
	l := len(node.Children)
	for i := range node.Children {
		idx := l - i - 1
		newnode.Children = append(newnode.Children, invert(node.Children[idx]))
	}
	return newnode
}
