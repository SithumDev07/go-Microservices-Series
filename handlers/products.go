package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SithumDev07/microservice/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
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
func (p *Products) addProduct(rw http.ResponseWriter, req *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

// * UPDATE | PUT

func (p* Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(res, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handling PUT request", id)
	product := &data.Product{}
	err = product.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, product)

	if err == data.ErrorProductNotFound {
		http.Error(res, "Product Not Found", http.StatusNotFound)
		return
	}
	
	if err != nil {
		http.Error(res, "Product Not Found", http.StatusInternalServerError)
		return
	}
}