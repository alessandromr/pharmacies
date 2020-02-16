package config

var (
	//RefreshRateHours indicate how often data is refreshed from PharmaciesDataSource
	RefreshRateHours = 24

	//PharmaciesDataSource indicate the url of the main external datasource for pharmacies
	PharmaciesDataSource = "https://dati.regione.campania.it/catalogo/resources/Elenco-Farmacie.geojson"
)
