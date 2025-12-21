package repository

import (
	"Invoice-Payment-System/internal/domain"
	"Invoice-Payment-System/internal/infrastructure/database"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepo struct {
	db         *database.MongoDB
	collection *mongo.Collection
}

func NewPaymentRepository(db *database.MongoDB) domain.PaymentRepository {
	collection := db.GetCollection("payments")
	return &PaymentRepo{db: db, collection: collection}
}

func (pr *PaymentRepo) LogPayment(payment *domain.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if payment.ID.IsZero() {
		payment.ID = primitive.NewObjectID()
	}

	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	_, err := pr.collection.InsertOne(ctx, payment)

	return err
}

func (pr *PaymentRepo) FindByReference(ref string) (*domain.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var payment domain.Payment
	err := pr.collection.FindOne(ctx, bson.M{"reference": ref}).Decode(&payment)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &payment, nil

}

func (pr *PaymentRepo) UpdateStatus(ref string, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"reference": ref}
	update := bson.M{"$set": bson.M{"status": status, "updatedAt": time.Now()}}

	_, err := pr.collection.UpdateOne(ctx, filter, update)

	return err
}

func (pr *PaymentRepo) FindByInvoiceID(invoiceID primitive.ObjectID) ([]*domain.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := pr.collection.Find(ctx, bson.M{"invoiceID": invoiceID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []*domain.Payment
	for cursor.Next(ctx) {
		var payment domain.Payment
		if err := cursor.Decode(&payment); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil

}
