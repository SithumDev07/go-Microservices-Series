package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SithumDev07/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	productList := data.GetProducts()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(data)
}