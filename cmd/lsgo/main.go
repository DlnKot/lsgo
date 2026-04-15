package main

import (
	"flag"
	"fmt"
	"lsgo/internal/app"
	"lsgo/internal/config"
	"lsgo/internal/fs"
	"os"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

func main() {
	var pathFlag string

	fileManager := fs.NewManager()

	flag.StringVar(&pathFlag, "path", fileManager.Pwd, "Use for set custom dir --path=[OPTION]")
	flag.Parse()

	cfg := config.New("")
	cfg.LoadConfig()

	height := 24
	if _, h, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		height = h
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	entries := fileManager.ReadDir(pathFlag)

	appInstance := app.NewApp(fileManager, &cfg)
	if err := appInstance.Run(entries, pathFlag, 0, 0, height); err != nil {
		fmt.Println("Error running app:", err)
	}
}
