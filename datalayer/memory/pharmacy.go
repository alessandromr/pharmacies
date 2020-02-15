package memory

import (
	"github.com/alessandromr/pharmacy/model"
	"sync"
)

//PharmaciesMemory implements the Pharmacies interface using local in memory persistance
type PharmaciesMemory struct {
	Pharmacies []model.Pharmacy
}

//GetPharmacies return a list of pharmacies
func (p *PharmaciesMemory) GetPharmacies() ([]model.Pharmacy, error) {
	return p.Pharmacies, nil
}

//SetPharmacies set a list of pharmacies, overwriting existing data
func (p *PharmaciesMemory) SetPharmacies(newList []model.Pharmacy) error {
	mutex := &sync.Mutex{}
	mutex.Lock()
	p.Pharmacies = newList
	mutex.Unlock()
	return nil
}
