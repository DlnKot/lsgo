package fs

import (
	"fmt"
	"os"
)

type Manager struct {
	Pwd string
}

func NewManager() *Manager {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "./"
	}

	return &Manager{
		Pwd: pwd,
	}
}

func (f Manager) ReadDir(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error for read dir", path, err)
	}

	return entries
}
