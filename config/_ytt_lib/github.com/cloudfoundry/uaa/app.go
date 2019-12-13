package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("[uaa] Request received")
	fmt.Fprintf(w, "<h1>[uaa] Hello %s!</h1>", os.Getenv("HELLO_MSG"))
	// fmt.Fprintf(w, "<p>local change</p>")
}

func main() {
	log.Print("[uaa] Server started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
