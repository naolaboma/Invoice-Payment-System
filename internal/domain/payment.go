package domain

import "time"

type Payment struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	InvoiceID string    `bson:"invoicID" json:"invoicID"`
	Reference string    `bson:"reference" json:"reference"`
	Amount    float64   `bson:"amount" json:"amount"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
