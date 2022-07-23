package main

import (
	"log"
	"net/http"

	"github.com/AssassinAsh/going-go/pkg/handlers"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)

	log.Println("Listening on port: ", portNumber)

	http.ListenAndServe(portNumber, nil)
}
