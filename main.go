package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/")
		time.Sleep(time.Duration(rand.Int()%3000) * time.Millisecond)
		fmt.Fprintf(w, "hello world")
	})

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/check")
		fmt.Fprintf(w, "ok")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
