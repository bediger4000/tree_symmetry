// Package multitree represents number valued nodes that form a tree
// with each node having an arbitrary number of child nodes
package multitree

// Node comprises elements of trees, each with any number of children
type Node struct {
	Data     int
	Children []*Node
}
