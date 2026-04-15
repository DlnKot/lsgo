package main

import (
	"flag"
	"fmt"
	"lsgo/internal"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

func main() {

	var pathFlag string
	fileManager := internal.NewFileManager()
	printer := internal.Printer{}

	flag.StringVar(&pathFlag, "path", fileManager.Pwd, "Use for set custom dir --path=[OPTION]")
	flag.Parse()

	config := internal.NewConfig("")
	config.LoadConfig()

	index := 0
	offset := 0

	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	entries := fileManager.ReadDir(pathFlag)

	for {
		ClearScreen()

		visible := limit(entries, height-2, offset)
		fmt.Println("Path:", pathFlag)
		for i, entry := range visible {
			printer.Print(i+offset, entry, index)
		}

		event, _, _ := keyboard.GetKey()

		switch event {
		case 'k':
			index--
		case 'j':
			index++
		case 'l':
			currentDir := entries[index]
			absPath := filepath.Join(pathFlag, currentDir.Name())

			if currentDir.IsDir() {
				entries = fileManager.ReadDir(absPath)
				pathFlag = absPath
				index = 0
				offset = 0
				continue
			} else {
				cmd := exec.Command(config.StandartEditor, absPath)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
				cmd.Run()
			}
		case 'h':
			parent := filepath.Dir(pathFlag)
			entries = fileManager.ReadDir(parent)
			pathFlag = parent
			index = 0
			offset = 0
		case 'q':
			return
		}

		if index < 0 {
			index = 0
		}

		if len(entries) == 0 {
			index = 0
		} else if index >= len(entries) {
			index = len(entries) - 1
		}

		if index >= offset+height-2 {
			offset++
		}

		if index < offset {
			offset--
		}
	}

}

func limit(entries []os.DirEntry, height int, offset int) []os.DirEntry {
	if offset > len(entries) {
		return []os.DirEntry{}
	}

	end := offset + height
	if end > len(entries) {
		end = len(entries)
	}

	return entries[offset:end]
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
