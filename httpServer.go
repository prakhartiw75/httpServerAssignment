package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	ID   int64  `json:id`
}

func ReqeustHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	dataReturn := Data{
		Name: "Prakhar Tiwari",
		ID:   26,
	}
	dataAccept := Data{}
	switch req.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(dataReturn)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
		}
	case "POST":
		err := json.NewDecoder(req.Body).Decode(&dataAccept)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
		}
		fmt.Fprintf(w, "The input recieved is=>%+v", dataAccept)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", ReqeustHandler)
	http.ListenAndServe(":8080", nil)
}
