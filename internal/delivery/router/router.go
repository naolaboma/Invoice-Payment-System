package router

import (
	"Invoice-Payment-System/internal/delivery/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(invoiceHandler *controllers.InvoiceHandler, paymentHandler *controllers.PaymentHandler) *gin.Engine{
	router := gin.Default()
	
	v1 := router.Group("/api/v1")
	{
		invoices := v1.Group("/invoices")
		{
			invoices.POST("/", invoiceHandler.CreateInvoice)
			invoices.GET("/", invoiceHandler.GetInvoicesBySender)
			invoices.GET("/:id", invoiceHandler.GetInvoiceByID)
		}
		payments := v1.Group("/payments")
		{
			payments.POST("/callback", paymentHandler.HandleCallBack)
			payments.GET("/:ref", paymentHandler.GetPaymentByReference)
		}
	}
	return router
}