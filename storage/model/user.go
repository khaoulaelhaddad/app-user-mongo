package model

type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}
