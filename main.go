package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my pretty alright site!</h1>")
}

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handlerFunc)
	//http.ListenAndServe(":3000", mux)

	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", nil)
}
