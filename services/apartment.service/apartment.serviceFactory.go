package apartment_service

import (
	"fmt"
	m "piso-scrapper/models"
)

type apartmentService interface {
	Create(apartment m.Apartment) error
	Read() (m.Apartments, error)
	FindCoincidences(apartment m.Apartment) bool
	UpdateCount(i int) error
	Count() int64
}

func GetApartmentFactory(collection string) (apartmentService, error) {
	switch collection {
	case "idealista":
		return &idealista{}, nil
	case "fotocasa":
		return &fotocasa{}, nil
	default:
		return nil, fmt.Errorf("wrong collection passed")
	}
}
