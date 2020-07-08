package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

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

//Validate the product structure
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
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

//DeleteProduct delete a product from the list
func DeleteProduct(id int) error {
	index := findIndexById(id)
	if index == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:index], productList[index+1])

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

func findIndexById(id int) int {
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
