package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alessandromr/pharmacy/datalayer/memory"
	"github.com/alessandromr/pharmacy/internal/rpcserver"
	"github.com/alessandromr/pharmacy/script"
)

func main() {
	pharmaciesDL := memory.PharmaciesMemory{}

	go script.SyncData(&pharmaciesDL)
	time.Sleep(time.Second * 2)

	// test := []float64{41.051814, 14.331329}

	// pharmacies, _ := pharmaciesDL.GetPharmacies()

	// pharmacyDistance := make([]PharmacyDistance, len(pharmacies))

	// for k, v := range pharmacies {
	// 	a := test[0] - v.Position.Latitude
	// 	b := test[1] - v.Position.Longitude
	// 	distance := math.Sqrt(math.Pow(a, float64(2)) + math.Pow(b, float64(2)))
	// 	pharmacyDistance[k] = PharmacyDistance{
	// 		Pharmacy: v,
	// 		Distance: distance,
	// 	}
	// }
	// sort.Slice(pharmacyDistance[:], func(i, j int) bool {
	// 	return pharmacyDistance[i].Distance < pharmacyDistance[j].Distance
	// })

	r := rpcserver.NewJsonRpcServer()
	log.Fatal(http.ListenAndServe(":8080", r))
}
