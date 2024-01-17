package main

import (
	"joe-l-mathew/Tambula-Ticket-Generator/config"
	"joe-l-mathew/Tambula-Ticket-Generator/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/api/generate_ticket", handlers.GenerateTicket).Methods("GET")
	r.HandleFunc("/api/games", handlers.GetAllTickets).Methods("GET")
	port := "8080"
	log.Printf("Server listening on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
