package main

import (
	"ascii-web/ascii"
	"net/http"
)

func main() {
	http.HandleFunc("/", ascii.Handler)

	http.ListenAndServe(":8080", nil)
}
