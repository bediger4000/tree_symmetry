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
		if symmetric(root) {
			return
		}
		os.Exit(1)
	}

	multitree.Print(root)
	fmt.Println()

	phrase := " "
	if !symmetric(root) {
		phrase = "not "
	}

	fmt.Printf("Tree is %ssymmetric\n", phrase)
}

type NodePair struct {
	node1 *multitree.Node
	node2 *multitree.Node

	next *NodePair
}

type NodePairStack struct {
	top *NodePair
}

func (nps *NodePairStack) Empty() bool {
	if nps.top == nil {
		return true
	}
	return false
}

func (nps *NodePairStack) Push(n1, n2 *multitree.Node) {
	pair := &NodePair{node1: n1, node2: n2}
	pair.next = nps.top
	nps.top = pair
}

func (nps *NodePairStack) Pop() (*multitree.Node, *multitree.Node) {
	top := nps.top
	nps.top = top.next
	return top.node1, top.node2
}

func symmetric(root *multitree.Node) bool {
	if root == nil {
		return true
	}

	stack := new(NodePairStack)
	stack.Push(root, root)

	for !stack.Empty() {
		node1, node2 := stack.Pop()
		if !compareNodes(node1, node2) {
			return false
		}
		// node1, node2 non-nil, have identical Data values,
		// same number of child nodes, at least
		l := len(node1.Children)
		for idx := range node1.Children {
			mirrorIdx := l - 1 - idx
			if mirrorIdx < idx {
				break
			}
			stack.Push(node1.Children[idx], node2.Children[mirrorIdx])
		}
	}

	return true
}

// compareNodes does comparison of nodes, not children
// of nodes. Easy stuff that would clutter the caller's code.
func compareNodes(node1, node2 *multitree.Node) bool {
	if node1 == nil {
		if node2 == nil {
			return true
		}
		return false
	}
	// node1 is non-nil from here
	if node2 == nil {
		return false
	}

	// node2 is non-nil from here

	if len(node1.Children) != len(node2.Children) {
		return false
	}

	if node1.Data != node2.Data {
		return false
	}

	return true
}
