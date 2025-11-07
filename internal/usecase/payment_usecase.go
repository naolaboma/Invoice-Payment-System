package usecase

import (
	"Invoice-Payment-System/internal/domain"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentUseCase struct{
	InvoiceRepo domain.InvoiceRepository
	PaymentRepo domain.PaymentRepository
}

func NewPaymentUseCase(invoicerepo domain.InvoiceRepository, paymentrepo domain.PaymentRepository) domain.PaymentUsecase{
	return &PaymentUseCase{
		InvoiceRepo: invoicerepo,
		PaymentRepo: paymentrepo,
	}
}

func (uc *PaymentUseCase) HandleCallBack(payment *domain.Payment) error{
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	
	err := uc.PaymentRepo.LogPayment(payment)
	if err != nil{
		return err
	}
	
	switch payment.Status {
	case "SUCCESS":
		return uc.InvoiceRepo.UpdateStatus(payment.InvoiceID, "PAID")
	case "FAILED":
		return uc.InvoiceRepo.UpdateStatus(payment.InvoiceID, "FAILED")
	default:
		// do nothing for pending or unknown statuses
		return nil
	}
}

func (uc *PaymentUseCase) GetPaymentByReference(ref string) (*domain.Payment, error){
	return uc.PaymentRepo.FindByReference(ref)
}

func (uc *PaymentUseCase) GetPaymentsByInvoice(invoiceID string) ([]*domain.Payment, error){
	objectID, err := primitive.ObjectIDFromHex(invoiceID)
	if err != nil {
		return nil, fmt.Errorf("invalid invoice ID: %v", err)
	}
	return uc.PaymentRepo.FindByInvoiceID(objectID)
}