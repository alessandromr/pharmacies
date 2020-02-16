package memory

import (
	"testing"

	"github.com/alessandromr/pharmacy/model"
)

func TestPharmaciesMemory_CompleteTest(t *testing.T) {
	pharmacies := []model.Pharmacy{}
	pMemory := PharmaciesMemory{}

	//Empty pharmacies
	pMemory.SetPharmacies(pharmacies)

	//Assert pharmacies are empty
	pharmacies, err := pMemory.GetPharmacies()
	if err != nil{
		t.Fatal(err)
	}
	if len(pharmacies) != 0{
		t.Fatal("Memory is not empty")
	}


	//Insert test data into pharmacies
	pharmacies = []model.Pharmacy{
		model.Pharmacy{
			Name: "Test Pharmacy 1",
			Position: model.Location{
				Latitude: 10.05615616,
				Longitude: 10.05615616,
			},
		},
		model.Pharmacy{
			Name: "Test Pharmacy 2",
			Position: model.Location{
				Latitude: 10.05615616,
				Longitude: 10.05615616,
			},
		},
	}
	pMemory.SetPharmacies(pharmacies)

	//Assert correct number of pharmacies is present
	pharmacies, err = pMemory.GetPharmacies()
	if err != nil{
		t.Fatal(err)
	}
	if len(pharmacies) != 2{
		t.Fatal("Cannot retrieve pharmacies from memory")
	}

}
