package mongo

import (
	"context"
	"url-shortener/internal/domains"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoURLRepository struct {
	collection *mongo.Collection
}

func NewMongoURLRepository(db *mongo.Database) *MongoURLRepository {
	return &MongoURLRepository{
		collection: db.Collection("urls"),
	}
}

func (r *MongoURLRepository) Save(url *domains.URL) error {
	_, err := r.collection.InsertOne(context.TODO(), url)
	return err
}

func (r *MongoURLRepository) FindByID(id string) (*domains.URL, error) {
	var url domains.URL
	if err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&url); err != nil {
		return nil, err
	}
	return &url, nil
}
