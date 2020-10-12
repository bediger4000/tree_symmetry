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
			fmt.Printf("node1 nil, node 2 (%p/%d) non-nil\n", node2, node2.Data)
			return false
		}
	}
	if node2 == nil {
		fmt.Printf("node 1 (%p/%d) non-nil, node 2 nil\n", node1, node1.Data)
		return false
	}

	// node1 and node2 not nil after this

	if node1.Data != node2.Data {
		fmt.Printf("node1 data %d != node2 data %d\n", node1.Data, node2.Data)
		return false
	}

	childCount := len(node1.Children)
	children := node1.Children
	mirrors := node2.Children
	ln2 := len(node2.Children)
	if ln2 > childCount {
		childCount = ln2
		children = node2.Children
		mirrors = node1.Children
	}

	for idx := range node1.Children {
		idxMirror := childCount - idx - 1

		if idxMirror < idx {
			break
		}

		if !symmetric(children[idx], mirrors[idxMirror]) {
			return false
		}
	}
	return true
}
