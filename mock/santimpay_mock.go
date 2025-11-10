package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type PaymentRequest struct {
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Reference string  `json:"reference"`
	Email     string  `json:"email"`
	Callback  string  `json:"callback"`
}

type PaymentResponse struct {
	PaymentURL string `json:"payment_url"`
	Reference  string `json:"reference"`
	Status     string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	// POST /v1/payments -> simulate creating a payment
	mux.HandleFunc("/v1/payments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req PaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		log.Printf("✅ Received mock payment request: %+v", req)

		// Generate a mock payment link
		paymentURL := fmt.Sprintf("http://localhost:8081/mock-payment/%s", req.Reference)

		resp := PaymentResponse{
			PaymentURL: paymentURL,
			Reference:  req.Reference,
			Status:     "PENDING",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// GET /mock-payment/:reference -> simulate payer visiting payment page
	mux.HandleFunc("/mock-payment/", func(w http.ResponseWriter, r *http.Request) {
		ref := r.URL.Path[len("/mock-payment/"):]
		status := "SUCCESS" // random success/fail

		fmt.Fprintf(w, `
			<html>
			<body>
				<h2>Mock SantimPay Payment Page</h2>
				<p>Reference: %s</p>
				<p>Amount: 100 ETB</p>
				<p>Status: %s</p>
				<p><b>Callback triggered to:</b> http://localhost:8080/api/v1/payments/callback</p>
			</body>
			</html>
		`, ref, status)

		// Simulate callback after 3 seconds
		go func() {
			time.Sleep(3 * time.Second)

			callbackBody := map[string]interface{}{
				"invoice_id": ref, // this is critical
				"reference":  ref,
				"status":     status,
				"amount":     100.0, // match your Payment struct
				"currency":   "ETB", // match your Payment struct
			}
			bodyBytes, _ := json.Marshal(callbackBody)

			resp, err := http.Post(
				"http://localhost:8080/api/v1/payments/callback",
				"application/json",
				bytes.NewReader(bodyBytes),
			)
			if err != nil {
				log.Printf("❌ Callback failed for ref=%s: %v", ref, err)
				return
			}
			defer resp.Body.Close()

			log.Printf("📡 Callback sent for ref=%s, status=%s, HTTP %d", ref, status, resp.StatusCode)
		}()
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	log.Println("🚀 Mock SantimPay server running on http://localhost:8081")
	log.Fatal(server.ListenAndServe())
}
