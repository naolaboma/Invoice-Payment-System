package main

import (
	"Invoice-Payment-System/internal/delivery/controllers"
	"Invoice-Payment-System/internal/delivery/router"
	"Invoice-Payment-System/internal/infrastructure/database"
	"Invoice-Payment-System/internal/infrastructure/gateway"
	"Invoice-Payment-System/internal/repository"
	"Invoice-Payment-System/internal/usecase"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	fmt.Println("MONGO_URI:", mongoURI)
	dbName := os.Getenv("DB_NAME")

	db, err := database.NewMongoDB(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Close()

	paymentGateway := gateway.NewSantimPayGateway()

	invoiceRepo := repository.NewInvoiceRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)

	invoiceUsecase := usecase.NewInvoiceUseCase(invoiceRepo, paymentGateway)
	paymentUsecase := usecase.NewPaymentUseCase(invoiceRepo, paymentRepo)

	invoiceHandler := controllers.NewInvoiceController(invoiceUsecase)
	paymentHandler := controllers.NewPaymentHandler(paymentUsecase)

	r := router.SetupRouter(invoiceHandler, paymentHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
