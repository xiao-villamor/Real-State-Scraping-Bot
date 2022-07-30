package Apartment_repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"piso-scrapper/database"
	m "piso-scrapper/models"
)

var collectionI = database.GetCollection("idealista")
var collectionF = database.GetCollection("fotocasa")

var ctx = context.Background()

func CreateI(appartment m.Apartment) error {
	var err error

	_, err = collectionI.InsertOne(ctx, appartment)

	if err != nil {
		return err
	}

	return nil

}

func ReadI() (m.Apartments, error) {

	var apartments m.Apartments

	filter := bson.D{}

	var err error
	var cur *mongo.Cursor

	cur, err = collectionI.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var apartment m.Apartment
		err = cur.Decode(&apartment)

		if err != nil {
			return nil, err
		}
		apartments = append(apartments, &apartment)
	}

	return apartments, nil
}

func UpdateI(cnt int) error {
	var err error

	filter := bson.M{"id": "count"}

	update := bson.M{
		"$set": bson.M{
			"price": cnt,
		},
	}

	_, err = collectionI.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil

}

func CountI() (int64, error) {

	var err error
	var estCount int64

	estCount, err = collectionI.EstimatedDocumentCount(context.TODO())

	if err != nil {
		return 0, err
	}
	return estCount, nil

}

func CreateF(appartment m.Apartment) error {
	var err error

	_, err = collectionF.InsertOne(ctx, appartment)

	if err != nil {
		return err
	}

	return nil

}

func ReadF() (m.Apartments, error) {

	var apartments m.Apartments

	filter := bson.D{}

	var err error
	var cur *mongo.Cursor

	cur, err = collectionF.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var apartment m.Apartment
		err = cur.Decode(&apartment)

		if err != nil {
			return nil, err
		}
		apartments = append(apartments, &apartment)
	}

	return apartments, nil
}

func UpdateF(cnt int) error {
	var err error

	filter := bson.M{"id": "count"}

	update := bson.M{
		"$set": bson.M{
			"price": cnt,
		},
	}

	_, err = collectionF.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil

}

func CountF() (int64, error) {

	var err error
	var estCount int64

	estCount, err = collectionF.EstimatedDocumentCount(context.TODO())

	if err != nil {
		return 0, err
	}
	return estCount, nil

}
