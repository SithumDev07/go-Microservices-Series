package handlers

import (
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

func (p*Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	
	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(rw, req)
		return
	}

	// catch All

	rw.WriteHeader(http.StatusMethodNotAllowed)
}


// * GET
func (p *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
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

	p.l.Printf("Product: %#v", product)
}