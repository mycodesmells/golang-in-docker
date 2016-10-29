package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello!"))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
