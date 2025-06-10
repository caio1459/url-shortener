package repositories

import (
	"context"
	"url-shortener/internal/domains"
	"url-shortener/internal/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type userRepository struct {
	collection *mongo.Collection
}

// Construtor que retorna a interface
func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

// FindByEmail implements interfaces.UserRepository.
func (u *userRepository) FindByEmail(email string) (*domains.User, error) {
	var user domains.User
	if err := u.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Save implements interfaces.UserRepository.
func (u *userRepository) Save(user *domains.User) error {
	_, err := u.collection.InsertOne(context.TODO(), user)
	return err
}
