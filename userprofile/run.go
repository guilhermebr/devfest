package main

import (
	"fmt"
	"net/http"

	"github.com/guilhermebr/devfest/userprofile/server"
)

func main() {

	server.RegisterHandlers()

	fmt.Println("Rodando em 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
