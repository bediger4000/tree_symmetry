package main

import (
	"flag"
	"fmt"
	"log"
	"tree_symmetry/multitree"
)

func main() {
	graphViz := flag.Bool("g", false, "produce GraphViz dot output on stdout")
	dfudsout := flag.Bool("D", false, "print DFUDS string representation on stdout")
	dfudsin := flag.Bool("d", false, "parse DFUDS string representation from command line")
	flag.Parse()

	var root *multitree.Node
	var err error

	if *dfudsin {
		root, err = multitree.ParseDFUDS(flag.Arg(0))
	} else {
		root, err = multitree.FromString(flag.Arg(0))
	}
	if err != nil {
		log.Fatal(err)
	}

	if *graphViz {
		multitree.Draw(root)
		return
	}

	if *dfudsout {
		multitree.PrintDFUDS(root)
		return
	}

	// default output, balanced parens on stdout
	multitree.Print(root)
	fmt.Println()
}
