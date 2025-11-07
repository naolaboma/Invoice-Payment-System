package usecase

import (
	"Invoice-Payment-System/internal/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceUseCase struct {
	InvoiceRepo domain.InvoiceRepository
}

func NewInvoiceUseCase(repo domain.InvoiceRepository) domain.InvoiceUsecase {
	return &InvoiceUseCase{
		InvoiceRepo: repo,
	}
}

func (uc *InvoiceUseCase) CreateInvoice(invoice *domain.Invoice) error {
	if invoice.Amount <= 0 {
		return errors.New("invoice amount must be positive")
	}
	invoice.Status = "PENDING"
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	invoice.ID = primitive.NewObjectID()
	return uc.InvoiceRepo.Create(invoice)
}

func (uc *InvoiceUseCase) GetInvoice(id string) (*domain.Invoice, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid invoice ID format")
	}
	return uc.InvoiceRepo.FindById(objectID)
}

func (uc *InvoiceUseCase) UpdateInvoiceStatus(id string, status string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid invoice ID format")
	}
	return uc.InvoiceRepo.UpdateStatus(objectID, status)
}

func (uc *InvoiceUseCase) GetInvoicesBySender(email string) ([]*domain.Invoice, error) {
	invoices, err := uc.InvoiceRepo.FindBySenderEmail(email)
	if err != nil {
		return nil, err
	}
	return invoices, nil

}
