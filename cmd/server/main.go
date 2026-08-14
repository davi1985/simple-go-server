package main

import (
	"fmt"
	"go-server/internal/handler"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handler.Form)
	http.HandleFunc("/hello", handler.Hello)

	fmt.Printf("Server running at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
