package jsonRpc

type Args struct {
	CurrentLocation Location
	Range int
	Limit int
}

type Location struct{
	Latitude float64
	Longitude float64
}

type Pharmacy struct{}

func (s *Pharmacy) SearchNearest(args *Args, reply *int) error {
	return nil
}
