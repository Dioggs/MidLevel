package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	/*
		ler arquivo de configuração
		lidar com qualquer tipo de request	
		passar request para algoritimo de validação
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	const port int = 3000

	fmt.Printf("\nListening on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
