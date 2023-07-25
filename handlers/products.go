package handlers

import (
	"coffee_tym/products"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Key struct{}

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
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

func (p *ProductsHandler) PostProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(Key{}).(products.Product)

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

func (p ProductsHandler) ProdValMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := products.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "failed to unmarshall json", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			if err := prod.Validate(); err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
		}

		req := r.WithContext(
			context.WithValue(
				r.Context(),
				Key{},
				prod))

		next.ServeHTTP(rw, req)
	})
}
