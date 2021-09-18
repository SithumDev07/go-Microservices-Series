package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	server := &http.Server{
		Addr: ":8081",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func ()  {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChanel := make(chan os.Signal)
	signal.Notify(signalChanel, os.Interrupt)
	signal.Notify(signalChanel, os.Kill)

	sig := <- signalChanel
	l.Println("Recieved terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	server.Shutdown(timeoutContext)
}