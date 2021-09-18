package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if req.Method == http.MethodPut {
		// Expect the id from the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(req.URL.Path, -1)
		
		if len(group) != 1 {
			p.l.Println("Invalid URI More than one ID")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		
		if len(group[0]) != 2 {
			p.l.Println("Invalid URI More than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		
		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		
		if err != nil {
			p.l.Println("Unable to convert to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		
		p.updateProduct(id, rw, req)
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
	data.AddProduct(product)
}

// * UPDATE | PUT

func (p* Products) updateProduct(id int, res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handling PUT request")
	product := &data.Product{}
	err := product.FromJSON(req.Body)
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