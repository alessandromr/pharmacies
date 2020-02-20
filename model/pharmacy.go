package model

type Pharmacy struct {
	Name     string
	Position Location
	Distance int
}

type PharmacyDistance struct {
	Pharmacy Pharmacy
	Distance int
}
