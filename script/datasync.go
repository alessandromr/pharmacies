package script

import (
	"github.com/alessandromr/pharmacy/adapter/web"
	"github.com/alessandromr/pharmacy/config"
	"github.com/alessandromr/pharmacy/datalayer"
	"log"
	"time"
)

//SyncData update pharmacies list every 24 hours
func SyncData(pharmaciesDL datalayer.Pharmacy) {
	updatePharmacies(pharmaciesDL)
	for range time.Tick(time.Hour * time.Duration(config.RefreshRateHours)) {
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
