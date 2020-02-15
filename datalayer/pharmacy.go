package datalayer

import "github.com/alessandromr/pharmacy/model"

type Pharmacy interface {
	GetPharmacies() ([]model.Pharmacy, error)
	SetPharmacies([]model.Pharmacy) error
}
