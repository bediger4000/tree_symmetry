package multitree

import (
	"log"
	"math"
	"strconv"
	"unicode"
)

// FromString parses its input string, turning it into a multitree
func FromString(str string) *Node {
	runes := []rune(str)
	root, _ := FromRunes(runes, 0, len(str))
	return root
}

// FromRunes parses its input []rune, turning it into a mulitree,
// filling in any child nodes described by the input []rune,
// returning the multitree and the offset of runes just after the
// parsed expression.
func FromRunes(runes []rune, offset int, end int) (*Node, int) {
	// Check that it has leading '(' and trailing '('
	if runes[0] != '(' || runes[end-1] != ')' {
		log.Printf("FromRunes(%q, %d, %d) not called with parens\n", string(runes), offset, end)
	}
	// remove leading '('
	offset++

	// read number's digits
	makenode, number, newoffset := readDigits(runes, offset, end)

	if !makenode {
		return nil, newoffset + 1
	}

	// create Node with number as Data
	node := &Node{Data: number}

	offset = newoffset

	// Loop over remainder of runes to find child nodes
	for offset < end-1 {
		// 1. eat whitespace
		offset = eatWhiteSpace(runes, offset, end)

		// 2. Find next matching ')'
		endoffset := findRightParen(runes, offset)

		// 3. Call FromRunes with this slice of runes
		childNode, xoffset := FromRunes(runes, offset, endoffset)

		// 4. Add new child node to array of Children, even if nil
		node.Children = append(node.Children, childNode)

		// what's relationship between endoffset and xoffset?
		offset = xoffset
	}

	// remove trailing ')'
	offset++

	return node, offset
}

// readDigits returns a number, converted from the leading digits
// of runes []rune, and an index of where the number ends.
func readDigits(runes []rune, offset int, end int) (bool, int, int) {
	var valueRunes []rune
	for {
		if runes[offset] == '(' {
			break
		}
		if runes[offset] == ')' {
			break
		}
		if offset == end {
			break
		}
		if unicode.IsSpace(runes[offset]) {
			break
		}
		valueRunes = append(valueRunes, runes[offset])
		offset++
	}

	if len(valueRunes) == 0 {
		return false, math.MinInt64, offset
	}

	number, err := strconv.Atoi(string(valueRunes))
	if err != nil {
		log.Print(err)
	}
	return true, number, offset
}

// findRightParen takes an array of runes, where '(' is at
// index 0, and a matching ')' is at some greater index.
// Returns that greater index
func findRightParen(r []rune, offset int) int {
	stack := make([]rune, 1)
	stack[0] = r[offset]
	end := offset

	for idx := offset + 1; len(stack) > 0; idx++ {
		switch r[idx] {
		case '(':
			stack = append(stack, '(')
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			}
		}
		end++
	}
	return end + 1
}

// eatWhiteSpace starts at index offset in runes,
// and returns the index of the next non-whitespace rune.
func eatWhiteSpace(runes []rune, offset int, end int) int {
	for unicode.IsSpace(runes[offset]) {
		offset++
		if offset == end {
			break
		}
	}
	return offset
}
