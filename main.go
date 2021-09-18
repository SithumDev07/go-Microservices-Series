package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SithumDev07/microservice/handlers"
	"github.com/gorilla/mux"
)

func main () {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Handlers
	productHandler := handlers.NewProducts(l)

	serverMux := mux.NewRouter()

	// * Routers
	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serverMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)
	
	postRouter := serverMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	server := &http.Server{
		Addr: ":8081",
		Handler: serverMux,
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