package main

import (
	"fmt"
	"log"
	"os"
	"tree_symmetry/multitree"
)

func main() {
	graphViz := false
	n := 1
	if os.Args[1] == "-g" {
		graphViz = true
		n = 2
	}
	root, err := multitree.FromString(os.Args[n])
	if err != nil {
		log.Fatal(err)
	}

	if graphViz {
		multitree.Draw(root)
		return
	}

	multitree.Print(root)
	fmt.Println()
}
