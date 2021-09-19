package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Sithum",
		Price: 1.36,
		SKU: "abc-avg-asf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}