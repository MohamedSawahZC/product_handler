package main

import (
	"fmt"
	"net/http"

	"Api/productApi"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/product/all", productApi.FindAll).Methods("GET")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
