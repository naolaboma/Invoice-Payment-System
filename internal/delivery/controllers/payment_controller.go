package controller

import (
	"Invoice-Payment-System/internal/domain"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentHandler struct{
	PaymentUseCase domain.PaymentUsecase
	Validate *validator.Validate
}

func NewPaymentHandler(paymentUsecase domain.PaymentUsecase) *PaymentHandler{
	return &PaymentHandler{
		PaymentUseCase: paymentUsecase,
		Validate: validator.New(),
	}
}

func (h *PaymentHandler) HandleCallBack(c *gin.Context) {
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.Validate.Struct(payment); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	
	err := h.PaymentUseCase.HandleCallBack(&payment)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}

func (h *PaymentHandler) GetPaymentByReference(c *gin.Context){
	ref := c.Param("ref")
	payment, err := h.PaymentUseCase.GetPaymentByReference(ref)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if payment == nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Payment not found"})
		return
	}
	
	c.JSON(http.StatusOK, payment)
}