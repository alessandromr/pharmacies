package logic

import (
	"sort"

	"github.com/alessandromr/pharmacy/datalayer/memory"
	"github.com/alessandromr/pharmacy/internal/coordinates"
	"github.com/alessandromr/pharmacy/model"
)

//SearchNearestPharmacy return a list of nearest pharmacies
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

	var pharmaciesInRange []model.PharmacyDistance
	//Get pharmacies in range
	for _, v := range pharmacyDistance {
		if v.Distance < rangeMt {
			pharmaciesInRange = append(pharmaciesInRange, v)
		}
	}

	//Sort pharmacies by distance
	sort.Slice(pharmaciesInRange[:], func(i, j int) bool {
		return pharmaciesInRange[i].Distance < pharmaciesInRange[j].Distance
	})

	if limit >= len(pharmaciesInRange) {
		return pharmaciesInRange, nil
	}
	return pharmaciesInRange[:limit], nil
}
