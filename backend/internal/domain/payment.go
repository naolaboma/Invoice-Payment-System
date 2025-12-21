package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	InvoiceID primitive.ObjectID `bson:"invoiceID" json:"invoiceID"`
	Reference string             `bson:"reference" json:"reference"`
	Amount    float64            `bson:"amount" json:"amount"`
	Link      string             `bson:"link,omitempty" json:"link,omitempty"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

//interfaces

type PaymentRepository interface {
	LogPayment(payment *Payment) error
	FindByReference(ref string) (*Payment, error)
	FindByInvoiceID(invoiceID primitive.ObjectID) ([]*Payment, error)
	UpdateStatus(ref string, status string) error
}

type PaymentUsecase interface {
	HandleCallBack(payment *Payment) error
	GetPaymentByReference(ref string) (*Payment, error)
	GetPaymentsByInvoice(invoiceID string) ([]*Payment, error)
}