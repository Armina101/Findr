package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thebravebyte/findr/internals"
)

// Todo: Implement the methods of the LinkerStore interface here that has to do with the students collection.

func (fd *FindrDB) CreateAccount(std internals.Students) (bool, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var res bson.M

	filter := bson.D{{Key: "email", Value: user.Email}}
	err := FindrData(fd.DB, colType).FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			std.Id = primitive.NewObjectID().Hex()
			_, err := NHIAData(cd.DB, colType).InsertOne(ctx, user)
			if err != nil {
				cd.Lg.Error.Println(err)
				return false, false, errors.New("cannot registered this account")
			}
			return true, false, nil
		}
		cd.Lg.Error.Fatalln(err)
	}
	return true, true, nil
}
