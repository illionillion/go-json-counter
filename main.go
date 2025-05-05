package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/illionillion/go-json-counter/server"
)

func main() {
	http.HandleFunc("/", server.CounterHandler)
	http.HandleFunc("/user/{name}", server.CounterHandler)
	fmt.Println("Server Start Up........")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("PORT: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
