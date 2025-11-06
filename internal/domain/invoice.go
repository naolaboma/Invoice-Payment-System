package domain

import (
	"context"
	"time"
)

type Invoice struct {
	ID           string    `bson:"_id,omitempty" json:"id"`
	SenderEmail  string    `bson:"senderEmail" json:"senderEmail"`
	PayerEmail   string    `bson:"payerEmail" json:"payerEmail"`
	Amount       float64   `bson:"amount" json:"amount"`
	Currency     string    `bson:"currency" json:"currency"`
	Description  string    `bson:"description" json:"description"`
	Status       string    `bson:"status" json:"status"`
	SantimPayRef *string   `bson:"santimPayRef,omitempty" json:"santimPayRef,omitempty"`
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
}

// interfaces
type InvoiceRepository interface {
	Create(ctx context.Context, invoice *Invoice) error
	FindById(ctx context.Context, id string) (*Invoice, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}

type InvoiceUsecase interface {
	CreateInvoice(ctx context.Context, invoice *Invoice) error
	GetInvoice(ctx context.Context, id string) (*Invoice, error)
	UpdateInvoice(ctx context.Context, id string, status string) error
}
