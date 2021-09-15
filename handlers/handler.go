package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

func GetHandler(w http.ResponseWriter, req *http.Request) {

	v := response{
		Name:    "Something",
		Success: true,
	}

	jsonData, _ := json.Marshal(v)
	fmt.Println(jsonData)

	fmt.Fprint(w, string(jsonData))
	fmt.Println("Request arrived")

}

//function display data

type dis struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Domain  string `json: "domain"`
}

func display(w http.ResponseWriter, req *http.Request) {
	d := dis{
		Name:    "Keshika Gupta",
		Address: "Noida",
		Age:     21,
		Domain:  "Golang",
	}
	jsonData, _ := json.Marshal(d)
	fmt.Println(jsonData)

	fmt.Fprint(w, string(jsonData))
	fmt.Println("Request arrived")

}
