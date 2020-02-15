package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alessandromr/pharmacy/adapter/web"
	"github.com/alessandromr/pharmacy/datalayer"
	"github.com/alessandromr/pharmacy/datalayer/memory"
	"github.com/alessandromr/pharmacy/presentation/jsonRpc"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func main() {
	pharmaciesDL := memory.PharmaciesMemory{}

	// go syncData(&pharmaciesDL)
	syncData(&pharmaciesDL)

	log.Fatal()

	//Create RPC Server
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")

	//Register Objects
	s.RegisterService(new(jsonRpc.Pharmacy), "")

	//Start HTTP Server
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.Handle("/jsonrpc/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//syncData update pharmacies list every 24 hours 
func syncData(pharmaciesDL datalayer.Pharmacy) {
	updatePharmacies(pharmaciesDL)
	for range time.Tick(time.Hour * 24) {
		updatePharmacies(pharmaciesDL)
	}
}

func updatePharmacies(pharmaciesDL datalayer.Pharmacy) {
	log.Println("Updating Pharmacies List")
	pharmacies, err := web.GetPharmacies()
	if err != nil {
		log.Println("Error retrieving Pharmacies List: ", err)
	}
	pharmaciesDL.SetPharmacies(pharmacies)
}
