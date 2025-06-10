package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func init() {
	log.Println("Initializing .env file...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Initializing MongoDB configuration...")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(os.Getenv("DATABASE"))
}

func GetMongoDB() *mongo.Database {
	if db == nil {
		log.Fatal("MongoDB is not initialized")
	}
	return db
}
