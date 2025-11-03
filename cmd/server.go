package main

import (
	"fmt"
	"net/http"
)

func serveDir(d *string, p *int) error {
	fs := http.FileServer(http.Dir(*d))

	http.Handle("/", noCache(fs))

	port, err := getPort(p)
	if err != nil {
		return err
	}

	fmt.Printf("Listening on http://localhost%s", port)
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
