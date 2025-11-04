package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		errAndExit(err, oldState)
	}

	// defer term.Restore(int(os.Stdin.Fd()), oldState)

	cwd, err := os.Getwd()
	if err != nil {
		errAndExit(err, oldState)
	}

	path := flag.String("path", cwd, "Path of directory to serve")
	port := flag.Int("port", 8080, "Port to serve on")
	flag.Parse()

	server, err := createServer(path, port)
	if err != nil {
		errAndExit(err, oldState)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errAndExit(err, oldState)
		}
	}()

	displayStatus(*path, server.Addr)

	go func() {
		for {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			if char == 'q' {
				exit(oldState, server)
			}
			fmt.Printf("You pressed: %q\r\n", char)
		}
	}()

	select {}
}

func exit(oldState *term.State, srv *http.Server) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9"))

	fmt.Println(style.Render("Shutting down server..."))

	srv.Close()

	term.Restore(int(os.Stdin.Fd()), oldState)
	os.Exit(0)
}

func errAndExit(e error, oldState *term.State) {
	term.Restore(int(os.Stdin.Fd()), oldState)

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("9"))

	fmt.Println(style.Render(e.Error()))
	os.Exit(1)
}
