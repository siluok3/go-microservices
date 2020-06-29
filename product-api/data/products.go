package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for the API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products define a slice of a Product structure
type Products []*Product

// ToJSON encodes the products response
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns all the available products
func GetProducts() Products {
	return productList
}

//Test entries for products
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Freddo Espresso",
		Description: "Greek, italian inspired, cold espresso",
		Price:       3.5,
		SKU:         "fr457",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Freddo Capuccino",
		Description: "Greek, italian inspired, cold espresso with milky foam",
		Price:       4.5,
		SKU:         "frca666",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
