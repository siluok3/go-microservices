package data

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for the API product
// swagger:model
type Product struct {
	// the id for this product
	//
	// required:true
	// min: 1
	ID int `json:"id"`
	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`
	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`
	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`
	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products define a slice of a Product structure
type Products []*Product

// GetProducts returns all the available products
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product by the id
func GetProductByID(id int) (*Product, error) {
	index := findIndexByProductID(id)

	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[index], nil
}

// AddProduct creates a new product
func AddProduct(p Product) {
	//get the next id
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

//UpdateProduct updates an existing product
func UpdateProduct(p Product) error {
	index := findIndexByProductID(p.ID)
	if index == -1 {
		return ErrProductNotFound
	}
	productList[index] = &p

	return nil
}

//DeleteProduct delete a product from the list
func DeleteProduct(id int) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:index], productList[index+1])

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func findIndexByProductID(id int) int {
	for index, product := range productList {
		if product.ID == id {
			return index
		}
	}

	return -1
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
