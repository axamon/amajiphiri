package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)
}

// HelloServer writes static response.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Amaji Phiri rocks!")
}
