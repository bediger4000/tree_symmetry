package multitree

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// FromString parses its input string, turning it into a multitree
func FromString(str string) *Node {
	runes := []rune(str)
	consumed, root, err := FromRunes(runes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem parsing %q: %v\n", str, err)
		return nil
	}
	if consumed != len(runes) {
		fmt.Fprintf(os.Stderr, "used %d of %d characters to construct tree\n",
			consumed, len(runes))
	}
	return root
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
