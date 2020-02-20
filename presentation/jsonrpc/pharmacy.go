package jsonrpc

import (
	"github.com/alessandromr/pharmacy/internal/logic"
	"github.com/alessandromr/pharmacy/model"
	"net/http"
)

type Pharmacy struct{}

type SearchNearestPharmacyParamas struct {
	CurrentLocation model.Location `json:"currentLocation"`
	Range           int            `json:"range"`
	Limit           int            `json:"limit"`
}

type SearchNearestPharmacyResponse struct {
	Pharmacies []model.PharmacyDistance `json:"pharmacies"`
}

func (s *Pharmacy) SearchNearestPharmacy(r *http.Request, args *SearchNearestPharmacyParamas, reply *SearchNearestPharmacyResponse) error {
	pharmacies, err := logic.SearchNearestPharmacy(args.CurrentLocation, args.Range, args.Limit)
	reply.Pharmacies = pharmacies
	return err
}
