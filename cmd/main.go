package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := flag.String("path", cwd, "Path of directory to serve")
	port := flag.Int("port", 8080, "Port to serve on")

	fmt.Printf("Serving %s\n", *path)

	if err := serveDir(path, port); err != nil {
		log.Fatal(err)
	}
}
