package main

// Check if a multiple-child tree is symmetric or not
// recursive version

import (
	"fmt"
	"log"
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
	root, err := multitree.FromString(os.Args[n])
	if err != nil {
		log.Fatal(err)
	}

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
	if node1 == nil && node2 == nil {
		return true
	}

	// at least one node is non-nil, but if
	// the other is nil, the tree isn't symmetrical
	if node1 == nil || node2 == nil {
		return false
	}

	// node1 and node2 not nil after this

	if node1.Data != node2.Data {
		return false
	}

	// node1 and node2 have the same number of children

	if len(node1.Children) != len(node2.Children) {
		return false
	}

	childCount := len(node1.Children)
	children := node1.Children
	mirrors := node2.Children

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
