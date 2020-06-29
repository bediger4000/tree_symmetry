# Daily Coding Problem: Problem #237


A tree is symmetric if its data and shape remain unchanged when it is
reflected about the root node. The following tree is an example:

```
        4
      / | \
    3   5   3
  /           \
 9             9
```

Given a k-ary tree, determine whether it is symmetric.

I take "k-ary" to mean that each node can have an arbitrary
number of child nodes.
An ordering of the child nodes exists.

## package multitree

I wrote a small Golang package to represent trees where nodes
have any number of children

Package multitree represents number valued nodes that form a tree
with each node having an arbitrary number of child nodes

```go
type Node struct {
	Data     int
	Children []*Node
}
```

## Algorithm

The symmetry check consists of a single recursive function:

```go
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
		if !symmetric(node1.Children[idx], node2.Children[idxMirror]) {
			return false
		}
	}
	return true
}
```
