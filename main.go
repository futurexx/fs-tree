package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

// CommandLineArgs - command line args structure
type CommandLineArgs struct {
	targetPath string
	maxDepth   uint
}

var parsedArgs CommandLineArgs

// CheckArgs - checking command line arguments
func CheckArgs() error {
	fileStat, err := os.Stat(parsedArgs.targetPath)
	if err != nil {
		errorString := fmt.Sprintf("Failed to match path `%s`\n", parsedArgs.targetPath)
		return errors.New(errorString)
	}

	if !fileStat.IsDir() {
		return errors.New("Target must be a dir")
	}

	return nil
}

func init() {
	var targetPath string
	var maxDepth uint

	flag.StringVar(&targetPath, "t", ".", "Entrypoint path")
	flag.UintVar(&maxDepth, "max-depth", 0, "Max depth of recursion (positive number)")

	flag.Parse()

	parsedArgs.targetPath = targetPath
	parsedArgs.maxDepth = maxDepth
}

func main() {
	err := CheckArgs()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsedArgs)
}
