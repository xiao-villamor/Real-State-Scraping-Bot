package apartment_service

import (
	"log"
	m "piso-scrapper/models"
	apartmentrepository "piso-scrapper/repositories/Apartment.repository"
)

type fotocasa struct {
}

func (f fotocasa) Create(apartment m.Apartment) error {
	err := apartmentrepository.CreateF(apartment)

	if err != nil {
		return err
	}

	return nil
}

func (f fotocasa) Read() (m.Apartments, error) {
	apartments, err := apartmentrepository.ReadF()

	if err != nil {
		return nil, err
	}

	return apartments, nil
}

func (f fotocasa) FindCoincidences(apartment m.Apartment) bool {
	var allApartments m.Apartments
	var err error

	allApartments, err = apartmentrepository.ReadF()
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

func (f fotocasa) UpdateCount(cnt int) error {
	err := apartmentrepository.UpdateF(cnt)
	if err != nil {
		return err
	}
	return nil
}

func (f fotocasa) Count() int64 {
	cnt, err := apartmentrepository.CountF()

	if err != nil {
		return 0
	}
	return cnt
}
