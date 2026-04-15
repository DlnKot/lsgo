package internal

import (
	"os"

	"github.com/fatih/color"
)

type Printer struct {
}

func (p Printer) Print(index int, entry os.DirEntry, selectedIndex int) {
	if index == selectedIndex {
		if entry.IsDir() {
			color.Blue("> %s/", entry.Name())
		} else {		
			color.Cyan("> %s", entry.Name())
		}
	} else {
		if entry.IsDir() {
			color.Blue("  %s/", entry.Name())
		} else {		
			color.Cyan("  %s", entry.Name())
		}
	}
}
