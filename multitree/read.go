package multitree

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// FromString parses its input string, turning it into a multitree
func FromString(str string) (*Node, error) {
	runes := []rune(str)
	consumed, root, err := FromRunes(runes)
	if err != nil {
		return nil, err
	}
	if consumed != len(runes) {
		err = fmt.Errorf("used %d of %d characters to construct tree\n",
			consumed, len(runes))
	}
	return root, err
}

// FromRunes parses its input []rune, turning it into a mulitree,
// filling in any child nodes described by the input []rune,
// returning the multitree and the count of runes consumed
func FromRunes(runes []rune) (int, *Node, error) {
	if runes[0] != '(' {
		return 0, nil, errors.New("first character not opening parenthesis")
	}

	foundClosing := false
	max := len(runes)
	consumed := 1 // skip opening parentheses
	var children []*Node
	var valueRunes []rune

loop:
	for consumed < max {
		switch runes[consumed] {
		case '(':
			c, n, e := FromRunes(runes[consumed:])
			consumed += c
			if e != nil {
				return consumed, nil, e
			}
			children = append(children, n)
		case ')':
			consumed++
			foundClosing = true
			break loop
		default:
			if unicode.IsSpace(runes[consumed]) {
				consumed++
				continue
			}
			valueRunes = append(valueRunes, runes[consumed])
			consumed++
		}
	}

	if !foundClosing {
		return consumed, nil, errors.New("failed to find closing paren")
	}

	if len(valueRunes) == 0 {
		// assume this is a nil node, even if len(children) > 0
		return consumed, nil, nil
	}

	datum, err := strconv.Atoi(string(valueRunes))
	if err != nil {
		return consumed, nil, err
	}

	newNode := &Node{
		Data:     datum,
		Children: children,
	}

	return consumed, newNode, nil
}

var leftParen = []byte{'('}
var rightParen = []byte{')'}

// PrintDFUDS write a Depth-first Unary Degree Sequence
// representation of a multitree on stdout
func PrintDFUDS(root *Node) {
	PrintfDFUDS(os.Stdout, root)
}

// PrintfDFUDS write a Depth-first Unary Degree Sequence
// representation of a multitree on an arbitrary io.Writer
// No newline. Provide your own.
func PrintfDFUDS(out io.Writer, root *Node) {
	out.Write(leftParen)
	dfuds(out, root)
}

func dfuds(out io.Writer, node *Node) {
	if node == nil {
		return
	}

	l := len(node.Children)

	for i := 0; i < l; i++ {
		out.Write(leftParen)
	}
	out.Write([]byte(strconv.Itoa(node.Data)))
	out.Write(rightParen)
	for _, child := range node.Children {
		dfuds(out, child)
	}
}

// ParseDFUDS parse Depth-first Unary Degree Sequence format
// string representations of multitrees.
func ParseDFUDS(stringrep string) (*Node, error) {
	runes := []rune(stringrep)

	// start at runes[1], there's an extra '(' at the beginning
	consumed, root, err := realParseDFUDS(runes[1:])
	if err != nil {
		return nil, err
	}

	if consumed != len(runes)-1 {
		return root, fmt.Errorf("consumed %d of %d characters\n", consumed, len(runes))
	}

	return root, nil
}

func realParseDFUDS(runes []rune) (int, *Node, error) {
	childCount := 0
	consumed := 0

	for runes[consumed] == '(' {
		consumed++
		childCount++
	}

	var valueRunes []rune

	for runes[consumed] != ')' {
		valueRunes = append(valueRunes, runes[consumed])
		consumed++
	}
	consumed++ // eat the ')'

	n, err := strconv.Atoi(strings.TrimSpace(string(valueRunes)))
	if err != nil {
		return consumed, nil, err
	}

	newnode := &Node{
		Data:     n,
		Children: make([]*Node, childCount, childCount),
	}

	for i := 0; i < childCount; i++ {
		c, n, e := realParseDFUDS(runes[consumed:])
		if e != nil {
			return consumed + c, nil, e
		}
		newnode.Children[i] = n
		consumed += c
	}

	return consumed, newnode, nil
}
