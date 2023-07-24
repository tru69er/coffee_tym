package handlers

import (
	"coffee_tym/products"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (p *ProductsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw)
		return
	}

	if r.Method == http.MethodPost {
		p.postProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.putProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *ProductsHandler) putProducts(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["id"]
	if query == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(query[0])
	if product := products.ProductList.Find(id); product == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	} else {
		var new products.Product
		err := new.FromJSON(r.Body)
		if err != nil {
			p.l.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		p.l.Println("New Product ", new)
		product.Update(new)
	}
}

func (p *ProductsHandler) postProduct(rw http.ResponseWriter, r *http.Request) {
	var prod products.Product
	err := prod.FromJSON(r.Body)

	if err != nil || !prod.Validate() {
		rw.WriteHeader(http.StatusBadRequest)
		p.l.Println(err)
		return
	}

	prod.ID = len(products.ProductList) + 1
	prod.CreatedOn = time.Now().UTC().String()
	prod.UpdatedOn = time.Now().UTC().String()
	prod.SKU = "tor baba"

	products.ProductList = append(products.ProductList, &prod)

	rw.WriteHeader(http.StatusOK)
}

func (p *ProductsHandler) getProducts(rw http.ResponseWriter) {
	lp := products.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		log.Fatal(err)
	}
}
