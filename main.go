package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func api_func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handle_http() {
	http.HandleFunc("/", api_func)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	resp, err := http.Get("http://example.com/")
	fmt.Fprint(os.Stdout, resp)
	fmt.Fprint(os.Stdout, err)
	handle_http()
}
