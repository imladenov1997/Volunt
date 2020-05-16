package db

import (
	"context"
	"fmt"
	"github.com/imladenov1997/volunt/graph/model" // temporary until the DB layer gets implemented
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GlobalExchange *model.Exchange
var MongoClient *mongo.Client
var MongoDB *mongo.Database

const databaseName = "Volunt"
const exchangeCollectionName = "exchanges"

type DB struct {}

func (db DB) CreateExchange(exchange *model.Exchange) error {
	exchangeCollection := MongoDB.Collection(exchangeCollectionName)
	db.createIndex(exchangeCollection)

	insertResult, err := exchangeCollection.InsertOne(context.TODO(), *exchange)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println(insertResult)
	}

	return nil
}

func (db DB) GetExchange(ID *string) *mongo.SingleResult {
	exchangeCollection := MongoDB.Collection(exchangeCollectionName)
	exchangeUndecoded := exchangeCollection.FindOne(context.TODO(), bson.M{
		"id": *ID,
	})

	return exchangeUndecoded
}

func (db DB) UpsertPersonToExchange(exchangeID *string, exchangePair *model.ExchangePair) error {
	person := exchangePair.Owner
	exchangeCollection := MongoDB.Collection(exchangeCollectionName)
	_, err := exchangeCollection.UpdateOne(context.TODO(), bson.M{
		"id": exchangeID,
	}, bson.M{
		"$set": bson.M{
			"people." + person.ID: exchangePair,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (db DB) UpdatePersonalBill(exchangeID *string, exchangePair *model.ExchangePair) error {
	personID := exchangePair.Owner.ID
	exchangeCollection := MongoDB.Collection(exchangeCollectionName)
	_, err := exchangeCollection.UpdateOne(context.TODO(), bson.M{
		"id": exchangeID,
	}, bson.M{
		"$set": bson.M{
			"people." + personID + ".fromvalue": exchangePair.FromValue,
			"people." + personID + ".tovalue": exchangePair.ToValue,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic("Cannot connect to Database")
	}

	MongoClient = client
	MongoDB = client.Database(databaseName)

	return client
}

func (db DB) createIndex(exchange *mongo.Collection) error {
	_, err := exchange.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"id": 1,
		},
		Options: options.Index().SetUnique(true),
	})

	return err

}
