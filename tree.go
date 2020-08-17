package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
)

// FileNode - node of fs tree
type FileNode struct {
	Parent  *FileNode
	Name    string
	AbsPath string
	IsDir   bool
	IsLast  bool
}

// TreeStat - stats about count of directories and files
type TreeStat struct {
	DirCount  uint
	FileCount uint
}

func getColorizedFileName(fileNode *FileNode) string {
	color := ColorReset
	if fileNode.IsDir {
		color = ColorBlue
	}
	return string(color) + fileNode.Name + string(ColorReset)
}

func getLine(fileNode *FileNode, level int) string {

	line := ""
	parent := fileNode.Parent
	for level-1 > 0 {
		line = IndentPrefix + line
		if !parent.IsLast {
			line = ParentPrefix + line
		}
		parent = parent.Parent
		level--
	}

	prefix := FilePrefix
	if fileNode.IsLast {
		prefix = LastFilePrefix
	}
	if fileNode.Parent == nil {
		prefix = ""
	}

	return fmt.Sprintf("%s%s%s\n", line, prefix, getColorizedFileName(fileNode))
}

func bufferingTree(buffer *bytes.Buffer, stat *TreeStat, fileNode *FileNode, level int) error {

	line := getLine(fileNode, level)
	buffer.WriteString(line)

	if fileNode.IsDir {
		stat.DirCount++
		files, err := ioutil.ReadDir(fileNode.AbsPath)
		if err != nil {
			return err
		}

		lastFileIndex := len(files) - 1
		for index, file := range files {
			isLast := index == lastFileIndex
			childFileNode := FileNode{
				Parent:  fileNode,
				Name:    file.Name(),
				AbsPath: path.Join(fileNode.AbsPath, file.Name()),
				IsDir:   file.IsDir(),
				IsLast:  isLast}
			err := bufferingTree(buffer, stat, &childFileNode, level+1)
			if err != nil {
				return err
			}
		}
	} else {
		stat.FileCount++
	}

	return nil
}

// PrintFilesTree - print files tree
func PrintFilesTree(args CLConfig) (bytes.Buffer, error) {
	buffer := bytes.Buffer{}
	_, fileName := path.Split(args.TargetPath)
	rootFile := FileNode{
		Parent:  nil,
		Name:    fileName,
		AbsPath: args.TargetPath,
		IsDir:   true,
		IsLast:  true}

	stat := TreeStat{}

	err := bufferingTree(&buffer, &stat, &rootFile, 0)

	statLine := fmt.Sprintf("\n%d directories, %d files", stat.DirCount, stat.FileCount)

	buffer.WriteString(statLine)

	return buffer, err
}
