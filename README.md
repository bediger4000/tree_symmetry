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
		if idxMirror < idx {break}
		if !symmetric(node1.Children[idx], node2.Children[idxMirror]) {
			return false
		}
	}
	return true
}
```

The function follows a node and its "reflection" through the tree.

The function has base cases for nil-value nodes,
one or the other "reflected" node not existing,
different data for node and reflected node,
different number of children for node and reflected node.

## Analysis

A web search turns up a lot of "are binary trees symmetric" solutions,
but very few multi-child tree solutions.

You have to assume some ordering of the child nodes of any given node,
otherwise the question doesn't make any sense -
it just becomes an exercize in combinatorics,
with the usual O(n!) or worse solution.

There are a few tricky parts.
The first is that instead of a recursive comparison having to
"go left" for one child, and "go right" for the other child,
the comparison has to loop over the child nodes of the current node,
comparing outermost children, next outermost children, etc.

The next tricky piece is to realize you don't have to loop through all of the child nodes,
just to the "middle" child node.
The comparison is some kind of left node to the mirror image right node.
The algorithm has already checked outermost left vs outermost right.
It doesn't need to check them the other way around.

The final tricky part is the "middle" child of a node with an odd number of children.
One wrong solution is to ignore middle children -
they are themselves if you reflect the tree,
but the middle subtree could be asymmetric.
Another wrong solution is to check symmetry of middle subtrees by themselves.
A tree could have different "middle" subtrees of outermost nodes.

One correct solution is to realize that the two nodes the recursive comparison
function tracks can both examine the same middle node.
The comparison function does not need a special case for middle children,
it just needs to compute which child node to follow so that the middle
child gets examined.

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
