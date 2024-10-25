// api/repository/user_repository.go

package repository

import (
	"api/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository interface for user operations
type UserRepository interface {
	// create new payload called user's payload
	FindUserByEmail(email string) (model.User, error)
	FindUserByMobile(mobile string) (model.User, error)
	CreateUser(user model.User) error
}

// userRepository struct that implements UserRepository
type userRepository struct {
	db *mongo.Client
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{db: db}
}

// FindUserByEmail retrieves a user by their email
func (r *userRepository) FindUserByEmail(email string) (model.User, error) {
	collection := r.db.Database("your_database_name").Collection("users") // Change to your database name
	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// FindUserByMobile retrieves a user by their mobile number
func (r *userRepository) FindUserByMobile(mobile string) (model.User, error) {
	collection := r.db.Database("your_database_name").Collection("users") // Change to your database name
	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"mobile_number": mobile}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// CreateUser adds a new user to the database
func (r *userRepository) CreateUser(user model.User) error {
	collection := r.db.Database("your_database_name").Collection("users") // Change to your database name
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}
