package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	URI      string
	Client   *mongo.Client
	Database *mongo.Database
}

//initialize new mongodb and connects to the database

func NewMongoDB(uri, invoice_payment_db string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB at %s: %w", uri, err)
	}

	// ping the database to verify the connection

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB at %s: %w", uri, err)
	}

	log.Printf("Successfully connected to MongoDB at %s, database: %s\n", uri, invoice_payment_db)

	return &MongoDB{
		URI:      uri,
		Client:   client,
		Database: client.Database(invoice_payment_db),
	}, nil
}

//return mongodb collection by name

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.Database.Collection(name)
}

//close the mongodb gracefully

func (m *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect mongodb client: %w", err)
	}

	log.Println("mongodb client disconnected successfully")
	return nil
}

// check the mongodb is alive

func (m *MongoDB) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("mongodb health check failed: %w", err)
	}
	log.Println("mongodb health check passed")
	return nil
}
