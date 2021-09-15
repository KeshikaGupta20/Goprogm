package main

import (
	"log"
	"net/http"

	"github.com/KeshikaGupta20/GoProgram/handlers"
)

func main() {

	http.HandleFunc("/getData", handlers.GetHandler)
	http.HandleFunc("/display", handlers.display)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error", err)
	}

}
