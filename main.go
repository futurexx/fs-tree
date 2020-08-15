package main

import (
	"flag"
	"fmt"
)

// CommandLineArgs - command line args structure
type CommandLineArgs struct {
	targetPath string
	maxDepth   int
}

var parsedArgs CommandLineArgs

func init() {
	var targetPath string
	var maxDepth int

	flag.StringVar(&targetPath, "t", ".", "Entrypoint path")
	flag.IntVar(&maxDepth, "max-depth", -1, "Max depth of recursion")

	flag.Parse()

	parsedArgs.targetPath = targetPath
	parsedArgs.maxDepth = maxDepth
}

func main() {
	fmt.Println(parsedArgs)
}
