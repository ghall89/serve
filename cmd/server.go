package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/lipgloss"
)

func createServer(dir *string, prt *int) (*http.Server, error) {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(*dir))
	mux.Handle("/", noCache(fs))

	port, err := getPort(prt)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	return srv, nil
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

func displayStatus(d string, p string) error {
	plainStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("3"))

	urlStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4")).
		Underline(true)

	hintStyle := lipgloss.NewStyle().
		PaddingTop(1).
		Faint(true)

	block := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			plainStyle.Render("Serving: "),
			urlStyle.Render(d),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			plainStyle.Render("Listening on: "),
			urlStyle.Render(fmt.Sprintf("http://localhost%s", p)),
		),
		hintStyle.Render("(q)uit"),
	)

	fmt.Println(block)
	return nil
}
