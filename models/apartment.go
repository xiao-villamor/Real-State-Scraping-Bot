package models

//Apartment List
type Apartment struct {
	Id        string `json:"id"`
	Direction string `json:"direction"`
	Price     string `json:"price"`
	Url       string `json:"url"`
}

type Apartments []*Apartment
