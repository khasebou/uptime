package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khasebou/uptime/handlers"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()
	router.HandleFunc("/metrics", handlers.ReceiveMetrics).Methods("POST") // Specify POST method

	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}
