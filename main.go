package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CLConfig - command line args structure
type CLConfig struct {
	TargetPath string
	MaxDepth   int
}

var parsedArgs CLConfig

// CheckArgs - checking command line arguments
func CheckArgs() error {
	fileStat, err := os.Stat(parsedArgs.TargetPath)
	if err != nil {
		errorString := fmt.Sprintf("Failed to match path `%s`\n", parsedArgs.TargetPath)
		return errors.New(errorString)
	}

	if !fileStat.IsDir() {
		return errors.New("Target must be a dir")
	}

	if parsedArgs.MaxDepth < 0 {
		return errors.New("`max-depth` param must be a positive")
	}

	return nil
}

func init() {
	var targetPath string
	var maxDepth int // TODO: support max-depth params

	flag.StringVar(&targetPath, "t", ".", "Entrypoint path")
	flag.IntVar(&maxDepth, "max-depth", 0, "Max depth of recursion (positive number)")

	flag.Parse()

	parsedArgs.TargetPath = targetPath
	parsedArgs.MaxDepth = maxDepth

	err := CheckArgs()
	if err != nil {
		log.Fatal(err)
	}

	absPath, err := filepath.Abs(parsedArgs.TargetPath)
	if err != nil {
		log.Fatal(err)
	}

	parsedArgs.TargetPath = absPath
}

func main() {
	treeBuffer, err := PrintFilesTree(parsedArgs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(treeBuffer.String())
}
