	package configuration

	import (
		"context"
		"log"
		"time"

		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)

	func ConnectDB() *mongo.Client {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = client.Ping(ctx, nil); err != nil {
			log.Fatalf("MongoDB connection error: %v", err)
		}

		log.Println("Connected to MongoDB")
		return client
	}
