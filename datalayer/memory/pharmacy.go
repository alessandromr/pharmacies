package memory

import (
	"github.com/alessandromr/pharmacy/model"
	"sync"
)

var pharmaciesMemory []model.Pharmacy

//PharmaciesMemory implements the Pharmacies interface using local in memory persistance
type PharmaciesMemory struct {
}

//GetPharmacies return a list of pharmacies
func (p *PharmaciesMemory) GetPharmacies() ([]model.Pharmacy, error) {
	return pharmaciesMemory, nil
}

//SetPharmacies set a list of pharmacies, overwriting existing data
func (p *PharmaciesMemory) SetPharmacies(newList []model.Pharmacy) error {
	mutex := &sync.Mutex{}
	mutex.Lock()
	pharmaciesMemory = newList
	mutex.Unlock()
	return nil
}
