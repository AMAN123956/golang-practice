package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test", testRouteHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func testRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "aman") // return text
}

type Response struct {
	Message string `json:"message"`
}

func testRouteHandler2(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "aman"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); 
	err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
