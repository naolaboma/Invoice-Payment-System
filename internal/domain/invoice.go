package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderEmail  string             `bson:"senderEmail" json:"senderEmail"`
	PayerEmail   string             `bson:"payerEmail" json:"payerEmail"`
	Amount       float64            `bson:"amount" json:"amount"`
	Currency     string             `bson:"currency" json:"currency"`
	Description  string             `bson:"description" json:"description"`
	Status       string             `bson:"status" json:"status"`
	SantimPayRef *string            `bson:"santimPayRef,omitempty" json:"santimPayRef,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// interfaces
type InvoiceRepository interface {
	Create(invoice *Invoice) error
	FindById(id primitive.ObjectID) (*Invoice, error)
	UpdateStatus(id primitive.ObjectID, status string) error
}

type InvoiceUsecase interface {
	CreateInvoice(invoice *Invoice) error
	GetInvoice(id string) (*Invoice, error)
	UpdateInvoice(id string, status string) error
}
