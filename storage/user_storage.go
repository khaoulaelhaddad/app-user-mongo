package storage

import (
	"context"
	"userApp/storage/dto"
	"userApp/storage/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	InsertUser(ctx context.Context, user dto.UserDto) (primitive.ObjectID, error)
	SelectUserByID(ctx context.Context, id primitive.ObjectID) (model.User, error)
	SelectUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, userDto dto.UserDto) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}
