package tree

import "fmt"

// NumericNode is an element of a binary tree with a numeric value
type NumericNode struct {
	Data  int
	Left  *NumericNode
	Right *NumericNode
}

func (n *NumericNode) String() string {
	return fmt.Sprintf("%d", n.Data)
}
