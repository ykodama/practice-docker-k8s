package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		log.Println("recieve request")
		fmt.Fprintf(w, "Hello Docker!!")
	})

	log.Println("start server")
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
