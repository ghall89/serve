package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		errAndExit(err, oldState, nil)
	}

	cwd, err := os.Getwd()
	if err != nil {
		errAndExit(err, oldState, nil)
	}

	path := flag.String("path", cwd, "Path of directory to serve")
	port := flag.Int("port", 8080, "Port to serve on")
	flag.Parse()

	server, err := createServer(path, port)
	if err != nil {
		errAndExit(err, oldState, server)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errAndExit(err, oldState, server)
		}
	}()

	displayStatus(*path, server.Addr)

	go func() {
		keyEvents, err := keyboard.GetKeys(10)
		if err != nil {
			panic(err)
		}

		for {
			event := <-keyEvents

			// ctrl-c
			if event.Rune == '\x00' && event.Key == 3 {
				exit(oldState, server)
			}

			switch event.Rune {
			case 'Q':
				fallthrough
			case 'q':
				exit(oldState, server)
			case 'O':
				fallthrough
			case 'o':
				address := fmt.Sprintf("http://localhost%s", server.Addr)
				exec.Command("open", address).Run()
			}
		}
	}()

	select {}
}

func exit(oldState *term.State, srv *http.Server) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9"))

	fmt.Println(style.Render("Shutting down server...\r"))

	srv.Close()

	term.Restore(int(os.Stdin.Fd()), oldState)
	os.Exit(0)
}

func errAndExit(e error, oldState *term.State, srv *http.Server) {
	term.Restore(int(os.Stdin.Fd()), oldState)

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("9"))

	fmt.Println(style.Render(e.Error(), "\r"))

	if srv != nil {
		srv.Close()
	}

	os.Exit(1)
}
