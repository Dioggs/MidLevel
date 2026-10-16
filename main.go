package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	const port int = 3000

	fmt.Printf("\nListening on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
