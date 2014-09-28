package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users/", ListUers).Methods("GET")
	r.HandleFunc("/users/", NewUser).Methods("POST")
	r.HandleFunc("/users/"+"{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/"+"{id}", UpdateUser).Methods("PUT")

	http.Handle("/users/", r)

	fmt.Println("Rodando em 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
