// Package classification of Product API
//
// Documentation for Product API
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SithumDev07/microservice/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// * GET
func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// * POST
func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {

	p.l.Println("Handling POST Request")

	product := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)
}

// * UPDATE | PUT
func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(res, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handling PUT request", id)
	product := req.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)

	if err == data.ErrorProductNotFound {
		http.Error(res, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Product Not Found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {
}

func (p Products) MiddlewareProductValidation(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		product := data.Product{}
		err := product.FromJSON(req.Body)
		if err != nil {
			p.l.Println("[Error] decentralizing product", err)
			http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		//Validate the product
		err = product.Validate()
		if err != nil {
			p.l.Println("[Error] validating product", err)
			http.Error(res, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		context := context.WithValue(req.Context(), KeyProduct{}, product)
		req = req.WithContext(context)
		nextHandler.ServeHTTP(res, req)
	})
}
