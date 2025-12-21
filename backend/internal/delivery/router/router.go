package router

import (
	"Invoice-Payment-System/internal/delivery/controllers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(invoiceHandler *controllers.InvoiceHandler, paymentHandler *controllers.PaymentHandler) *gin.Engine {
	router := gin.Default()

	// CORS Configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	router.GET("/success", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Payment successful!"})
	})

	router.GET("/failure", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Payment failed. Please try again."})
	})
	return router
}
