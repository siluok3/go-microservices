package data

import (
	"encoding/json"
	"fmt"
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

// FromJSON converts the request parameters
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)

}

// ToJSON encodes the products response
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns all the available products
func GetProducts() Products {
	return productList
}

// AddProduct creates a new product
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

//UpdateProduct updates an existing product
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p

	return nil
}

// ErrProductNotFound throws the explanation error
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
