package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/")
		fmt.Fprintf(w, "hello world")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("NOMAD_PORT_webapp"), nil))
}
