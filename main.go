package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Parse areguments
	var port int
	flag.IntVar(&port, "port", 8080, "port number for Yahtzee server")
	flag.Parse()
	// Create Yahtzee
	log.Println("Initializing Yahtzee...")
	y := Yahtzee{}
	y.Init()
	// Create API
	router := mux.NewRouter()
	router.HandleFunc("/playerName", y.getPlayerName).Methods("GET")
	router.HandleFunc("/playerName", y.setPlayerName).Methods("PUT")
	router.HandleFunc("/isYahtzee", y.isYahtzee).Methods("GET")
	router.HandleFunc("/rollDice", y.rollDice).Methods("POST")
	router.HandleFunc("/rollDie/{id}", y.rollDie).Methods("POST")
	router.HandleFunc("/dice", y.getDice).Methods("GET")
	router.HandleFunc("/die/{id}", y.getDie).Methods("GET")
	router.HandleFunc("/die", y.setDie).Methods("PUT")
	// Launch and listen
	log.Printf("Listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
