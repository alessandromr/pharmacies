package web

import (
	"testing"
)

func TestIntegration_GetPharmacies(t *testing.T) {
	pharmacies, err := GetPharmacies()
	if err != nil{
       t.Fatal(err)
	}

	if len(pharmacies) < 1 {
		t.Fatal("No pharmacies")
	}
}
