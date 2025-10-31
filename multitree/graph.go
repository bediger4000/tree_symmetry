package multitree

import (
	"fmt"
	"io"
	"os"
)

// GraphViz dot format graph on stdout
func Draw(node *Node) {
	Drawf(os.Stdout, node)
}

// GraphViz dot format graph on whatever io.Writer out is
func Drawf(out io.Writer, root *Node) {
	fmt.Fprintf(os.Stdout, "digraph g {\n")
	DrawPrefixed(os.Stdout, root, "N")
	fmt.Fprintf(os.Stdout, "\n}\n")
}

// DrawPrefixed outputs dot language directives for the current
// node, and the edges directed to its children.
func DrawPrefixed(out io.Writer, node *Node, prefix string) {
	fmt.Fprintf(out, "%s%p [label=\"%v\"];\n", prefix, node, node.Data)
	for i, child := range node.Children {
		if child == nil {
			fmt.Fprintf(out, "%s%p%d [shape=\"point\"];\n", prefix, node, i)
			fmt.Fprintf(out, "%s%p -> %s%p%d;\n", prefix, node, prefix, node, i)
			continue
		}
		DrawPrefixed(out, child, prefix)
		fmt.Fprintf(out, "%s%p -> %s%p;\n", prefix, node, prefix, child)
	}
}
