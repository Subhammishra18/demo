package model

type User struct {
	Name         string `bson:"name"`
	Email        string `bson:"email"`
	Age          int    `bson:"age"`
	Sex          string `bson:"sex"`
	Password     string `bson:"password"`
	MobileNumber string `bson:"mobile_number"` // Added for signup
}
