package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", ":80", "Port to use")

func main() {

	flag.Parse()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Printf("Listening on %s ...", *port)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
