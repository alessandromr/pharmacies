package web

import (
	"encoding/json"
	"github.com/alessandromr/pharmacy/config"
	"github.com/alessandromr/pharmacy/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type PharmacyResponses struct {
	Features []Feature
}

type Feature struct {
	Type       string
	Geometry   Geometry
	Properties Properties
}

type Geometry struct {
	Type        string
	Coordinates []float64
}

type Properties struct {
	Descrizione string
}

//GetPharmacies retrieve a list of pharmacies from web exposed catalog
func GetPharmacies() ([]model.Pharmacy, error) {
	var bodyBytes []byte
	var pharmacies []model.Pharmacy

	sleep := 1
	for {
		response, err := http.Get(config.PharmaciesDataSource)

		if err != nil {
			log.Printf("Updating Pharmacies List Failed, retrying after %d seconds\n", sleep)
			sleep++
			if sleep > 5 {
				return pharmacies, err
			}
			time.Sleep(time.Duration(sleep) * time.Second)
		} else {
			bodyBytes, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return pharmacies, err
			}
			break
		}
	}

	pharmaciesResponse := PharmacyResponses{}
	err := json.Unmarshal(bodyBytes, &pharmaciesResponse)
	if err != nil {
		return pharmacies, err
	}

	for _, v := range pharmaciesResponse.Features {
		pharmacy := model.Pharmacy{
			Name: v.Properties.Descrizione,
			Position: model.Location{
				Latitude:  v.Geometry.Coordinates[1],
				Longitude: v.Geometry.Coordinates[0],
			},
		}
		pharmacies = append(pharmacies, pharmacy)
	}

	return pharmacies, nil
}
