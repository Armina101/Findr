package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type FindrDB struct {
	DB *mongo.Client
}

// NewLinkerDB returns a new constructor instance of the CompliantDB struct.
func NewFindrDB(db *mongo.Client) *FindrDB {
	return &FindrDB{
		DB: db,
	}
}

// LinkerData returns a reference to a collection in the nhia database.
func FindrData(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("Findr").Collection(collectionName)
}
