/* package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
*/