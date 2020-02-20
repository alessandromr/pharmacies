package logic

import (
	"sort"

	"github.com/alessandromr/pharmacy/datalayer/memory"
	"github.com/alessandromr/pharmacy/internal/coordinates"
	"github.com/alessandromr/pharmacy/model"
)

func SearchNearestPharmacy(location model.Location, rangeMt int, limit int) ([]model.PharmacyDistance, error) {
	pharmaciesDL := memory.PharmaciesMemory{}

	//Get pharmacies list from datalayer
	pharmacies, _ := pharmaciesDL.GetPharmacies()
	pharmacyDistance := make([]model.PharmacyDistance, len(pharmacies))

	//Calculate distancies from pharmacies
	for k, v := range pharmacies {
		distance := coordinates.CalculateDistance(
			coordinates.Coordinates{
				Lat: location.Latitude,
				Lon: location.Longitude,
			},
			coordinates.Coordinates{
				Lat: v.Position.Latitude,
				Lon: v.Position.Longitude,
			},
		)

		pharmacyDistance[k] = model.PharmacyDistance{
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
		if v.Distance > rangeMt {
			pharmacyDistance = pharmacyDistance[:k]
			break
		}
	}

	if limit >= len(pharmacyDistance) {
		return pharmacyDistance, nil
	}
	return pharmacyDistance[:limit], nil
}
