package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("[cc] Request received")
	fmt.Fprintf(w, "<h1>[cc] Hello %s!</h1>", os.Getenv("HELLO_MSG"))

	resp, err := http.Get(os.Getenv("UAA_URL"))
	if err != nil {
		fmt.Fprintf(w, "[cc] error: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "[cc] error: %s", err)
		return
	}

	fmt.Fprintf(w, "[cc] uaa response: %s", body)
}

func main() {
	log.Print("[cc] Server started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
