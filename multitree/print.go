package multitree

import (
	"fmt"
	"io"
	"os"
)

// Print puts a text representation of a multitree,
// parseable by func FromString, on stdout
func Print(node *Node) {
	Printf(os.Stdout, node)
}

// Print puts a text representation of a multitree, parseable by
// func FromString, on whatever out io.Writer is.
func Printf(out io.Writer, node *Node) {
	if node == nil {
		fmt.Fprintf(out, "()")
		return
	}
	fmt.Fprintf(out, "(%v", node.Data)
	for _, child := range node.Children {
		fmt.Fprintf(out, " ")
		Printf(out, child)
	}
	fmt.Fprintf(out, ")")
}
