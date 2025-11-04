package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		errAndExit(err)
	}

	path := flag.String("path", cwd, "Path of directory to serve")
	port := flag.Int("port", 8080, "Port to serve on")
	flag.Parse()

	if err := serveDir(path, port); err != nil {
		errAndExit(err)
	}
}

func errAndExit(e error) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("9"))

	fmt.Println(style.Render(e.Error()))
	os.Exit(1)
}
