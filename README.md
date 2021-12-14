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

This has appeared as "Daily Coding Problem: Problem #1077 [Easy]"

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


## Analysis

A web search turns up a lot of "are binary trees symmetric" solutions,
but very few multi-child tree solutions.

If you assume that an ordering of child nodes exists,
you end up having to deal with cases like this:

```
tree 1: (4 (3 () (6)) (5) (3 (6) ()))
tree 2: (4 (3 () (6)) (5) (3 (6)   ))
```

I think tree1 is not symmetric.

Is tree 2 symmetric? 
The 3-valued nodes have different length child-node-arrays,
but the left-hand 3-valued node has a nil-valued node just
to keep the order correct.

There are a few tricky parts.
The first is that instead of a recursive comparison having to
"go left" for one child, and "go right" for the other child,


This is probably a decent interview problem.
Multi-child trees don't occur too often in day-to-day programming,
so candidates probably haven't worked with them before.
I know I had to build some infrastructure
(new Node type, input and output functions)
to do this problem.
My binary tree infrastructure was not suitable.

The interviewer would get to see if a candidate could ask good questions
since ordering of child nodes is potentially a sticky issue.
The interviewer could see if a candidate can do array indexing and recursive problem solving.

When careless mirror child node selection can lead to examining the tree twice,
and the idea that both comparison nodes can refer to the same node exist,
the interviewer might get to see if the candidate can demonstrate a little insight.

A [non-recursive solution](sym_iterative.go) might be given extra points.
It's basically the same as a depth-first traverse of a binary tree,
except that the code has to keep track of two nodes in the tree to compare.
An iterative solution would require more "infrastructure" code,
a stack or queue to keep track of pairs of nodes-to-compare,
but none of this is very special code.

Overall, a decent problem for an interview, I think.

## Knuth Transform

From Wikipedia's [Left-child right-sibling binary tree](https://en.wikipedia.org/wiki/Left-child_right-sibling_binary_tree):

```
The process of converting from a k-ary tree to an LC-RS binary tree is
sometimes called the Knuth transform. To form a binary tree from an
arbitrary k-ary tree by this method, the root of the original tree is made the
root of the binary tree. Then, starting with the root, each node's leftmost
child in the original tree is made its left child in the binary tree, and its
nearest sibling to the right in the original tree is made its right child in
the binary tree.
```

[Another explanation](https://xlinux.nist.gov/dads/HTML/binaryTreeRepofTree.html).

I wrote [code to do that transform](transform.go), just for fun.
This is the example tree shown in the Wikipedia article:

```
$ ./transform '(1 (2 (5) (6)) (3) (4 (7 (8)(9))))' > x.dot
$ dot -Tpng -o x.png
```

That will get you a PNG graphic showing a multitree and the equivalent binary
tree after a Knuth Transform.

The reverse is just as easy: the special format binary trees produced by a Knuth Transform
can be returned to a k-ary tree.

```
$ ./transform -r '(1 (2 (5) (6)) (3) (4 (7 (8)(9))))' > x.dot
$ dot -Tpng -o x.png
```

---

## Daily Coding Problem: Problem #686 [Hard]

This problem was asked by Adobe.

You are given a tree with an even number of nodes.
Consider each connection between a parent and child node to be an "edge".
You would like to remove some of these edges,
such that the disconnected subtrees that remain
each have an even number of nodes.

For example, suppose your input was the following tree:

```
   1
  / \ 
 2   3
    / \ 
   4   5
 / | \
6  7  8
```

In this case, removing the edge (3, 4) satisfies our requirement.

Write a function that returns the maximum number of edges you can remove
while still satisfying this requirement.

### Analysis

I haven't worked on this yet.
