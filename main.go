package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello World")
		d, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":8081", nil)
}