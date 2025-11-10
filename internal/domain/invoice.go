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
	PaymentLink  string             `bson:"paymentLink,omitempty" json:"payment_link,omitempty"`
	SantimPayRef *string            `bson:"santimPayRef,omitempty" json:"santimPayRef,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// interfaces
type InvoiceRepository interface {
	Create(invoice *Invoice) error
	FindById(id primitive.ObjectID) (*Invoice, error)
	UpdateStatus(id primitive.ObjectID, status string) error
	FindBySenderEmail(email string) ([]*Invoice, error)
	UpdatePaymentInfo(id primitive.ObjectID, paymentLink, reference string) error
}

type PaymentGateway interface {
	CreatePayment(invoice *Invoice) (paymentURL, reference string, err error)
}

type InvoiceUsecase interface {
	CreateInvoice(invoice *Invoice) (string, error)
	GetInvoice(id string) (*Invoice, error)
	UpdateInvoiceStatus(id string, status string) error
	GetInvoicesBySender(email string) ([]*Invoice, error)
}