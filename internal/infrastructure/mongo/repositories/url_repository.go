package repositories

import (
	"context"
	"url-shortener/internal/domains"
	"url-shortener/internal/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type urlRepository struct {
	collection *mongo.Collection
}

func NewMongoURLRepository(db *mongo.Database) interfaces.URLRepository {
	return &urlRepository{
		collection: db.Collection("urls"),
	}
}

// FindByID implements interfaces.URLRepository.
func (u *urlRepository) FindByID(id string) (*domains.URL, error) {
	var url domains.URL
	if err := u.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&url); err != nil {
		return nil, err
	}
	return &url, nil
}

// Save implements interfaces.URLRepository.
func (u *urlRepository) Save(url *domains.URL) error {
	_, err := u.collection.InsertOne(context.TODO(), url)
	return err
}
