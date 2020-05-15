package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "server up and running...")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	log.Println("server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
