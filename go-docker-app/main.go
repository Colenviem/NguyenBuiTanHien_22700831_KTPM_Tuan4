package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker Go!")
	})

	fmt.Println("Server đang chạy tại http://localhost:8082")
	http.ListenAndServe(":8080", nil)
}