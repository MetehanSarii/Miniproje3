package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "Kullanici ID  " + id

	message := API{messageConcat}
	output, err := json.Marshal(message)

	if err != nil {
		log.Println("hata oluştu")
	}
	fmt.Fprintf(w, string(output))
}
func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user/{id:[0-9]+}", hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
