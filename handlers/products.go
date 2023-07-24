package handlers

import (
	"coffee_tym/products"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (p *ProductsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		p.postProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *ProductsHandler) PutProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if prod := products.Find(id); prod == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	} else {
		var new products.Product
		err := new.FromJSON(r.Body)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		prod.Update(new)
	}
	rw.WriteHeader(http.StatusOK)
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

func (p *ProductsHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := products.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		log.Fatal(err)
	}
}
