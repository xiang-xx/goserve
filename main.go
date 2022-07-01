package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dir := flag.String("d", ".", "static file dir")
	port := flag.String("p", "8888", "port")
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	fmt.Printf("Listening on http://localhost:%s/index.html", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
