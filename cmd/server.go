package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/lipgloss"
)

func serveDir(d *string, p *int) error {
	fs := http.FileServer(http.Dir(*d))

	http.Handle("/", noCache(fs))

	port, err := getPort(p)
	if err != nil {
		return err
	}

	displayStatus(*d, port)

	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

func displayStatus(d string, p string) {
	plainStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("3"))

	urlStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4")).
		Underline(true)

	// hintStyle := lipgloss.NewStyle().
	// 	PaddingTop(1).
	// 	Faint(true)

	fmt.Println(lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			plainStyle.Render("Serving "),
			urlStyle.Render(d),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			plainStyle.Render("Listening on "),
			urlStyle.Render(fmt.Sprintf("http://localhost/%s", p)),
		),
		// hintStyle.Render("(q)uit, (o)pen in browser"),
	))
}
