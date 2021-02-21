package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// u, err := url.Parse(s)
	// key: test
	// value: wafrule
	keys, ok := r.URL.Query()["test"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'test' is missing if you want to evaluate query strings")
		fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
		return
	}
	fmt.Fprintf(w, "Hi there, I love %s!\n", keys[0])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
