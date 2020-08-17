package main

import (
	"bytes"
	"io/ioutil"
	"path"
)

// // FileNode - node of fs tree
// type FileNode struct {
// 	Name    string
// 	AbsPath string
// 	Depth   int
// 	IsDir   bool
// 	IsFirst bool
// 	IsLast  bool
// }

func getIndent(count int) string {
	tabulation := ""
	for i := 0; i < count*4; i++ {
		tabulation += " "
	}

	return tabulation
}

func bufferingTree(buffer *bytes.Buffer, filePath string, level int) error {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		stringToWrite := getIndent(level)
		buffer.WriteString(stringToWrite + file.Name() + "\n")

		if file.IsDir() {
			absPath := path.Join(filePath, file.Name())
			bufferingTree(buffer, absPath, level+1)
		}

	}

	return nil
}

// PrintFilesTree - print files tree
func PrintFilesTree(args CLConfig) (bytes.Buffer, error) {
	buffer := bytes.Buffer{}

	err := bufferingTree(&buffer, args.TargetPath, 0)

	return buffer, err
}
