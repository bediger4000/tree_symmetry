package main

// Check if a multiple-child tree is symmetric or not

import (
	"fmt"
	"os"
	"tree_symmetry/multitree"
)

func main() {
	viaExitStatus := false
	n := 1
	if os.Args[1] == "-q" {
		n = 2
		viaExitStatus = true
	}
	root := multitree.FromString(os.Args[n])

	if viaExitStatus {
		if symmetric(root, root) {
			return
		}
		os.Exit(1)
	}

	multitree.Print(root)
	fmt.Println()

	phrase := " "
	if !symmetric(root, root) {
		phrase = "not "
	}

	fmt.Printf("Tree is %ssymmetric\n", phrase)
}

func symmetric(node1, node2 *multitree.Node) bool {
	if node1 == nil {
		if node2 == nil {
			return true
		} else {
			return false
		}
	}
	if node2 == nil {
		return false
	}

	// node1 and node2 not nil after this

	if node1.Data != node2.Data {
		return false
	}

	ln1 := len(node1.Children)
	ln2 := len(node2.Children)
	if ln1 != ln2 {
		return false
	}

	for idx := range node1.Children {
		idxMirror := ln1 - idx - 1
		if idxMirror < idx {
			break
		}
		if !symmetric(node1.Children[idx], node2.Children[idxMirror]) {
			return false
		}
	}
	return true
}
