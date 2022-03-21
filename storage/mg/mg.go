package mg

import "go.mongodb.org/mongo-driver/mongo"

type DbConnexion struct {
	Db *mongo.Database
}
