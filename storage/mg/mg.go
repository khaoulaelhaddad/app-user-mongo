package mg

import (
	"userApp/internal"

	"go.mongodb.org/mongo-driver/mongo"
)

type DbConnexion struct {
	Db *mongo.Database
}

func (db *DbConnexion) GetCollection(nameCollection string) *mongo.Collection {
	return db.Db.Collection(nameCollection)
}

func (db *DbConnexion) GetUsersCollection() *mongo.Collection {
	return db.GetCollection(internal.User)
}
