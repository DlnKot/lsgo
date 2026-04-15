package internal

import (
	"fmt"
	"os"
)

type FileManager struct {
	Pwd string
}

func NewFileManager() *FileManager {

	pwd, err := os.Getwd()
	if err != nil {
		pwd = "./"
	}

	return &FileManager{
		Pwd: pwd,
	}
}

func (f FileManager) ReadDir(path string) []os.DirEntry  {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error for read dir", path, err)
	}

	return entries
}