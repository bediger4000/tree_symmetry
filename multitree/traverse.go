package multitree

import "fmt"

type Visitorfunc func(*Node)

func PrintVisitor(node *Node) {
	fmt.Printf("%d ", node.Data)
}

func PreorderVisitorTraverse(node *Node, fn Visitorfunc) {
	fn(node)
	for i := range node.Children {
		PreorderVisitorTraverse(node.Children[i], fn)
	}
}
