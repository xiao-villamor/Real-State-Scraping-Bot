package apartment_service

import (
	"log"
	m "piso-scrapper/models"
	apartmentrepository "piso-scrapper/repositories/Apartment.repository"
)

type idealista struct {
}

func (i idealista) Create(apartment m.Apartment) error {
	err := apartmentrepository.CreateI(apartment)

	if err != nil {
		return err
	}

	return nil
}

func (i idealista) Read() (m.Apartments, error) {
	apartments, err := apartmentrepository.ReadI()

	if err != nil {
		return nil, err
	}

	return apartments, nil
}

func (i idealista) FindCoincidences(apartment m.Apartment) bool {
	var allApartments m.Apartments
	var err error

	allApartments, err = apartmentrepository.ReadI()

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range allApartments {
		if value.Url == apartment.Url {
			return false
		}
	}
	return true
}

func (i idealista) UpdateCount(cnt int) error {
	err := apartmentrepository.UpdateI(cnt)
	if err != nil {
		return err
	}
	return nil
}

func (i idealista) Count() int64 {
	cnt, err := apartmentrepository.CountI()

	if err != nil {
		return 0
	}
	return cnt
}
