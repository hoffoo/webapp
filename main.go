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
		time.Sleep(rand.Int() % 3000 * time.Millisecond)
		fmt.Fprintf(w, "hello world")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
