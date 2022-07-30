package apartment_service

import (
	m "piso-scrapper/models"
	apartmentrepository "piso-scrapper/repositories/Apartment.repository"
	"testing"
)

func TestCreate(t *testing.T) {

	aptest := m.Apartment{
		Id:        "12",
		Direction: "testdir",
		Price:     "344",
		Url:       "https://",
	}

	err := apartmentrepository.CreateI(aptest)

	if err != nil {
		t.Error("Prueba creacion cagaste")
		t.Fail()
	} else {
		t.Log("turrou que flipas")
	}

}
