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
