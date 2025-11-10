package repository

import (
	"Invoice-Payment-System/internal/domain"
	"context"
	"errors"
	"fmt"
	"time"

	"Invoice-Payment-System/internal/infrastructure/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceRepo struct {
	db *database.MongoDB
	collection *mongo.Collection
}

func NewInvoiceRepository(db *database.MongoDB) domain.InvoiceRepository{
		collection := db.GetCollection("invoices")
		return &InvoiceRepo{db: db, collection: collection}
}

func (ir *InvoiceRepo) Create(invoice *domain.Invoice) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	
	_, err := ir.collection.InsertOne(ctx, invoice)
	if err != nil{
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("invoice with id already exists: %w", err)
		}
		if errors.Is(err, context.DeadlineExceeded){
			return fmt.Errorf("database context deadline exceeded: %w", err)
		}
		
		return fmt.Errorf("failed to create invoice: %w", err)
	}
	
	return nil
}

func (ir *InvoiceRepo) FindById(id primitive.ObjectID) (*domain.Invoice, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var invoice domain.Invoice
	filter := bson.M{"_id": id}
	err := ir.collection.FindOne(ctx, filter).Decode(&invoice)
	
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return nil, errors.New("invoice not found")
		}
		return nil, fmt.Errorf("database error in findById: %w", err)
	}
	return &invoice, nil
}

func(ir *InvoiceRepo) UpdateStatus(id primitive.ObjectID, status string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	update := bson.M{
		"$set": bson.M{
			"status": status,
			"updatedAt": time.Now(),
		},
	}
	
	opts := options.Update().SetUpsert(false)
	res, err := ir.collection.UpdateByID(ctx, id, update, opts)
	
	if err != nil{
		if errors.Is(err, context.DeadlineExceeded){
			return fmt.Errorf("database cotext deadline exceeded: %w", err)
		}
		return fmt.Errorf("failed to update invoice status: %w", err)
	}
	
	if res.MatchedCount == 0{
		return fmt.Errorf("invouce not found")
	}
	
	return nil
}

func (ir *InvoiceRepo) FindBySenderEmail(email string) ([]*domain.Invoice, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	cursor, err := ir.collection.Find(ctx, bson.M{"senderEmail": email})
	
	if err != nil{
		return nil, err
	}
	defer cursor.Close(ctx)
	
	var invoices []*domain.Invoice
	for cursor.Next(ctx){
		var invoice domain.Invoice
		if err := cursor.Decode(&invoice); err != nil{
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (ir *InvoiceRepo) UpdatePaymentInfo(id primitive.ObjectID, paymentLink, reference string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	update := bson.M{
		"$set": bson.M{
			"paymentLink": paymentLink,
			"santimPayRef": reference,
			"updatedAt": time.Now(),
		},
	}
	
	res, err := ir.collection.UpdateByID(ctx, id, update)
	
	if err != nil{
		return fmt.Errorf("failed to update payment info: %w", err)
	}
	
	if res.MatchedCount == 0{
		return fmt.Errorf("invoice not found")
	}
	return nil
}