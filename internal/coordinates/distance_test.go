package coordinates

import (
	"testing"
)

type CoordinateTestCase struct {
	Cordinates1      []float64 //Lat-Lon
	Cordinates2      []float64 //Lat-Lon
	ExpectedDistance int
}

func TestCalculateDistanceSamePoint(t *testing.T) {

	testCases := []CoordinateTestCase{
		CoordinateTestCase{
			Cordinates1:      []float64{41.10938993, 15.0324625},
			Cordinates2:      []float64{41.10938993, 15.0324625},
			ExpectedDistance: 0,
		},
		CoordinateTestCase{
			Cordinates1:      []float64{-12.099395, 17.912292},
			Cordinates2:      []float64{41.10938993, 15.0324625},
			ExpectedDistance: 5923945,
		},
		CoordinateTestCase{
			Cordinates1:      []float64{41.1093899, 15.032101},
			Cordinates2:      []float64{41.10938993, 15.0324625},
			ExpectedDistance: 30,
		},
		CoordinateTestCase{
			Cordinates1:      []float64{-12.099395, -16.912292},
			Cordinates2:      []float64{-13.099395, -176.912292},
			ExpectedDistance: 16449098,
		},
		CoordinateTestCase{
			Cordinates1:      []float64{-0.099395, -0.912292},
			Cordinates2:      []float64{179.099395, 179.912292},
			ExpectedDistance: 144120,
		},
	}

	for _, v := range testCases {

		distance := CalculateDistance(
			Coordinates{
				Lat: v.Cordinates1[0],
				Lon: v.Cordinates1[1],
			},
			Coordinates{
				Lat: v.Cordinates2[0],
				Lon: v.Cordinates2[1],
			},
		)

		if distance != v.ExpectedDistance {
			t.Fatal()
		}
	}
}
