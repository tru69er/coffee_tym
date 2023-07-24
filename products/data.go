package products

import (
	"encoding/json"
	"io"
	"log"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"description"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"updatedOn"`
	DeletedOn string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w) // faster than marshall
	return e.Encode(p)
	// enc := json.NewEncoder(rw)
	// enc.Encode(lp)
	/*
		if send, err := json.Marshal(lp); err != nil {
			p.l.Fatal(err)
			http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
		} else {
			rw.Write(send)
		}
	*/
}

func (p *Product) Update(new Product) {
	changed := false
	if new.Desc != "" && new.Desc != p.Desc {
		changed = true
		p.Desc = new.Desc
	}
	if new.Price != 0 && new.Price != p.Price {
		changed = true
		p.Price = new.Price
	}
	if new.Name != "" && new.Name != p.Name {
		changed = true
		p.Name = new.Name
	}

	if changed {
		p.UpdatedOn = time.Now().UTC().String()
	}
}

func Find(id int) *Product {
	if id <= len(ProductList) {
		return (ProductList)[id-1]
	}
	return nil
}

func (p *Product) Validate() bool {
	log.Println(p)
	return (p.Name != "" && p.Desc != "" && p.Price > 0)
}

func (p *Product) FromJSON(w io.Reader) error {
	d := json.NewDecoder(w)
	return d.Decode(p)
	/*
		buf, err := io.ReadAll(r.Body)

		if err != nil {
			p.l.Fatal(err)
		}

		data := products.Product{}
		err =json.Unmarshal(buf, &data)

		if err != nil {
			p.l.Fatal(err)
		}

		p.l.Println(data)
	*/
}

func GetProducts() Products {
	return ProductList
}

var ProductList = Products{
	&Product{
		ID:        1,
		Name:      "Latte",
		Desc:      "Frothy milky coffee",
		Price:     2.45,
		SKU:       "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Short and strong coffee without milk",
		Price:     1.99,
		SKU:       "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
