package jsonrpc

import (
	"math"
	"net/http"
	"sort"
	"time"

	"github.com/alessandromr/pharmacy/datalayer/memory"
	"github.com/alessandromr/pharmacy/internal/coordinates"
	"github.com/alessandromr/pharmacy/model"
	"github.com/alessandromr/pharmacy/script"
)

type Pharmacy struct{}

type SearchNearestPharmacyParamas struct {
	CurrentLocation Location `json:"currentLocation"`
	Range           int      `json:"range"`
	Limit           int      `json:"limit"`
}

type SearchNearestPharmacyResponse struct {
	Pharmacies []PharmacyDistance `json:"pharmacies"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PharmacyDistance struct {
	Pharmacy model.Pharmacy
	Distance int
}

func (s *Pharmacy) SearchNearestPharmacy(r *http.Request, args *SearchNearestPharmacyParamas, reply *SearchNearestPharmacyResponse) error {
	const earthRadiusKm = 6371
	pharmaciesDL := memory.PharmaciesMemory{}

	go script.SyncData(&pharmaciesDL)
	time.Sleep(time.Second * 2)

	//Get pharmacies list from datalayer
	pharmacies, _ := pharmaciesDL.GetPharmacies()
	pharmacyDistance := make([]PharmacyDistance, len(pharmacies))

	//Calculate distancies from pharmacies
	for k, v := range pharmacies {
		distance := coordinates.CalculateDistance(
			coordinates.Coordinates{
				Lat: args.CurrentLocation.Latitude,
				Lon: args.CurrentLocation.Longitude,
			},
			coordinates.Coordinates{
				Lat: v.Position.Latitude,
				Lon: v.Position.Longitude,
			},
		)

		pharmacyDistance[k] = PharmacyDistance{
			Pharmacy: v,
			Distance: distance,
		}
	}

	//Sort pharmacies by distance
	sort.Slice(pharmacyDistance[:], func(i, j int) bool {
		return pharmacyDistance[i].Distance < pharmacyDistance[j].Distance
	})

	//Remove pharmacies out of range
	for k, v := range pharmacyDistance {
		if v.Distance > args.Range {
			pharmacyDistance = pharmacyDistance[:k]
			break
		}
	}

	if args.Limit >= len(pharmacyDistance) {
		reply.Pharmacies = pharmacyDistance
		return nil
	}
	reply.Pharmacies = pharmacyDistance[:args.Limit]
	return nil
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
