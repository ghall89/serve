package main

import (
	"flag"
	"log"
)

func main() {
	path := flag.String("path", ".", "Path of directory to serve")
	port := flag.Int("port", 8080, "Port to serve on")

	err := serveDir(path, port)
	if err != nil {
		log.Fatal(err)
	}
}
