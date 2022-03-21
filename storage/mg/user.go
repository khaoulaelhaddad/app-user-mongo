package mg

import (
	"context"
	"errors"
	"userApp/storage/dto"
	"userApp/storage/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DbConnexion) InsertUser(ctx context.Context, user dto.UserDto) (id primitive.ObjectID, err error) {

	collection := db.Db.Collection("users")
	res, err := collection.InsertOne(ctx, bson.M{"name": user.Name, "age": user.Age})
	if err != nil {

		return id, err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		return id, nil
	}
	return primitive.ObjectID{}, errors.New("ERROR in convert id to objectID")

}

func (db *DbConnexion) SelectUserByID(ctx context.Context, id primitive.ObjectID) (user model.User, err error) {
	collection := db.Db.Collection("users")
	res := collection.FindOne(ctx, bson.M{"_id": id})

	if err != nil {

		return user, err
	}
	err = res.Decode(&user)

	return

}

func (db *DbConnexion) SelectUsers(ctx context.Context) (users []model.User, err error) {
	collection := db.Db.Collection("users")
	res, err := collection.Find(ctx, bson.D{})
	if err != nil {

		return users, err
	}
	err = res.Decode(&users)
	for res.Next(ctx) {
		user := model.User{}
		err = res.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return

}

func (db *DbConnexion) UpdateUser(ctx context.Context, id primitive.ObjectID, userDto dto.UserDto) (err error) {
	collection := db.Db.Collection("users")
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": userDto},
	)
	return

}

func (db *DbConnexion) DeleteUser(ctx context.Context, id primitive.ObjectID) (err error) {
	collection := db.Db.Collection("users")
	_, err = collection.DeleteMany(
		ctx,
		bson.M{"_id": id},
	)
	return

}
