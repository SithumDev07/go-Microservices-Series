package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SithumDev07/microservice/handlers"
)

func main () {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":8081", sm)
}