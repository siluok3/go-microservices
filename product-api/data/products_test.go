package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Kiri",
		Price: 5.00,
		SKU:   "kou-ra-da",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
