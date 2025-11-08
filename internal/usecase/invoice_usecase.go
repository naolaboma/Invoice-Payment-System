package usecase

import (
	"Invoice-Payment-System/internal/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceUseCase struct {
	InvoiceRepo domain.InvoiceRepository
	PaymentGate domain.PaymentGateway
}

func NewInvoiceUseCase(repo domain.InvoiceRepository, gateway domain.PaymentGateway) domain.InvoiceUsecase {
	return &InvoiceUseCase{
		InvoiceRepo: repo,
		PaymentGate: gateway,
	}
}

func (uc *InvoiceUseCase) CreateInvoice(invoice *domain.Invoice) (string, error) {
	if invoice.Amount <= 0 {
		return "", errors.New("invoice amount must be positive")
	}
	invoice.Status = "PENDING"
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	invoice.ID = primitive.NewObjectID()
	if err := uc.InvoiceRepo.Create(invoice); err != nil {
		return "", err
	}
	paymentURL, reference, err := uc.PaymentGate.CreatePayment(invoice)
	if err != nil {
		return "", err
	}

	invoice.PaymentLink = paymentURL
	invoice.SantimPayRef = &reference

	return paymentURL, nil
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
