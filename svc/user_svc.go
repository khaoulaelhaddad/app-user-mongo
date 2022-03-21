package svc

import (
	"context"
	"userApp/storage"
	"userApp/storage/dto"
	"userApp/storage/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(ctx context.Context, userStorage storage.User, user dto.UserDto) (res model.User, err error) {
	id, err := userStorage.InsertUser(ctx, user)
	if err != nil {
		return
	}
	res, err = userStorage.SelectUserByID(ctx, id)
	return
}

func SelectUserByID(ctx context.Context, userID primitive.ObjectID, userStorage storage.User) (res model.User, err error) {
	res, err = userStorage.SelectUserByID(ctx, userID)
	return
}

func SelectUsers(ctx context.Context, userStorage storage.User) (res []model.User, err error) {
	res, err = userStorage.SelectUsers(ctx)
	return
}

func UpdateUser(ctx context.Context, userID primitive.ObjectID, userDto dto.UserDto, userStorage storage.User) (err error) {
	err = userStorage.UpdateUser(ctx, userID, userDto)
	return
}

func DeleteUser(ctx context.Context, userID primitive.ObjectID, userStorage storage.User) (err error) {
	err = userStorage.DeleteUser(ctx, userID)
	return
}
