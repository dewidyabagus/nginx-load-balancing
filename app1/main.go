package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>From Service 1</h1>"))

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	})

	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8000)

	log.Println("HTTP Start From", address)

	http.ListenAndServe(address, nil)
}
