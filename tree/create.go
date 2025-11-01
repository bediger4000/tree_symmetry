package tree

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

// createNumericNode fills in a struct NumericNode
// from a string argument, then returns a pointer to that NumericNode,
// except as something that fits interface Node.
func createNumericNode(stringValue string) *NumericNode {
	n, err := strconv.Atoi(stringValue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating tree.Numeric node from %q: %v\n",
			stringValue, err)
		return nil
	}
	return &NumericNode{Data: n}
}

// CreateNumericFromString parses a single string
// like "(2(0()(12))(34(-2)(100)))"
// and turns it into a binary tree of the given shape.
func CreateNumericFromString(stringrep string) (*NumericNode, error) {
	runes := []rune(stringrep)
	origLength := len(runes)

	consumed, root, err := GeneralCreateFromString(runes)

	if consumed != origLength {
		return nil, fmt.Errorf("string rep of length %d, consumed only %d runes",
			origLength, consumed)
	}
	return root, err
}

// GeneralCreateFromString uses a func argument to create a tree
// of type Node. It returns the root Node, on which the caller should
// do a type assertion to get the correct type. The int return is the
// number of runes read from the front of the runes []rune argument.
func GeneralCreateFromString(runes []rune) (int, *NumericNode, error) {

	if runes[0] != '(' {
		return 0, nil, errors.New("first character not opening parenthesis")
	}

	var value []rune
	var left, right *NumericNode
	setLeft := false
	foundClosing := false

	max := len(runes)
	consumed := 1 // skip opening parentheses

loop:
	for consumed < max {

		switch runes[consumed] {
		case '(':
			c, n, e := GeneralCreateFromString(runes[consumed:])
			consumed += c
			if e != nil {
				return consumed, nil, e
			}
			if !setLeft {
				left = n
				setLeft = true
			} else {
				right = n
			}
		case ')':
			consumed++
			foundClosing = true
			break loop
		default:
			if unicode.IsSpace(runes[consumed]) {
				consumed++
				continue
			}
			value = append(value, runes[consumed])
			consumed++
		}
	}

	if !foundClosing {
		return consumed, nil, errors.New("failed to find closing paren")
	}

	if len(value) == 0 {
		if setLeft {
			return consumed, nil, errors.New("no data value with child node(s)")
		}
		return consumed, nil, nil
	}

	newNode := createNumericNode(string(value))
	newNode.Left = left
	newNode.Right = right

	return consumed, newNode, nil
}

// Printf emits lisp-like balanced parens representation
// of a tree on out.
func Printf(out io.Writer, node *NumericNode) {
	fmt.Fprintf(out, "(%s", node)

	leftNil := node.Left == nil
	rightNil := node.Right == nil

	if !leftNil && !rightNil {
		out.Write([]byte(" "))
		Printf(out, node.Left)
		out.Write([]byte(" "))
		Printf(out, node.Right)
	} else if !leftNil && rightNil {
		out.Write([]byte(" "))
		Printf(out, node.Left)
	} else if leftNil && !rightNil {
		out.Write([]byte(" () "))
		Printf(out, node.Right)
	} // else both child nodes empty
	out.Write([]byte(")"))
}
