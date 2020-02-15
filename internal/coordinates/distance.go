package coordinates

import (
	"math"
)

const (
	earthRadius = 6371 // radius of the earth
)

type Coordinate struct {
	Lat float64
	Lon float64
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func CalculateDistance(p1, p2 Coordinate) (km int) {
	lat1 := degreesToRadians(p1.Lat)
	lon1 := degreesToRadians(p1.Lon)
	lat2 := degreesToRadians(p2.Lat)
	lon2 := degreesToRadians(p2.Lon)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2) * 
	math.Pow(math.Sin(dLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return int(c * earthRadius * 1000)
}