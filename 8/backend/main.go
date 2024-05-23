package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Docker!")
	})

	fmt.Println("Starting server at port 3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		fmt.Println(err)
	}
}
