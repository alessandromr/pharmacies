package script

import (
	"testing"

	"github.com/alessandromr/pharmacy/datalayer/memory"
)

func TestIntegration_SyncData(t *testing.T) {
	pharmacyDL := memory.PharmaciesMemory{}

	pharmacies, err := pharmacyDL.GetPharmacies()
	if err != nil{
		t.Fatal(err)
	}
	if len(pharmacies) != 0{
		t.Fatal("Memory is not empty")
	}
	
	updatePharmacies(&pharmacyDL)

	pharmacies, err = pharmacyDL.GetPharmacies()
	if err != nil{
		t.Fatal(err)
	}
	if len(pharmacies) == 0{
		t.Fatal("Error retrieving pharmacies")
	}
}
