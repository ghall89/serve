package main

import (
	"fmt"
	"net/http"
)

func serveDir(d *string, p *int) error {
	var err error

	fs := http.FileServer(http.Dir(*d))
	http.Handle("/", fs)

	port, err := getPort(p)
	if err != nil {
		return err
	}

	fmt.Printf("Listening on http://localhost%s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}
