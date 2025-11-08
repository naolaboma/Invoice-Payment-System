package gateway

import (
	"Invoice-Payment-System/internal/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SantimPayGateway struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewSantimPayGateway() domain.PaymentGateway {
	return &SantimPayGateway{
		apiKey:  os.Getenv("SANTIMPAY_API_KEY"),
		baseURL: os.Getenv("SANTIMPAY_API_URL"),
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (g *SantimPayGateway) CreatePayment(invoice *domain.Invoice) (string, string, error) {
	payload := map[string]interface{}{
		"amount":    invoice.Amount,
		"currency":  "ETB",
		"reference": invoice.ID.Hex(),
		"email":     invoice.PayerEmail,
		"callback":  os.Getenv("CALLBACK_URL"),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal SantimPay payload: %v", err)
	}

	req, err := http.NewRequest("POST", g.baseURL+"/payments", bytes.NewBuffer(body))
	if err != nil {
		return "", "", fmt.Errorf("failed to create SantimPay request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.apiKey)

	resp, err := g.client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("SantimPay API call failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", "", fmt.Errorf("SantimPay returned status %v", resp.Status)
	}

	var response struct {
		PaymentURL string `json:"payment_url"`
		Reference  string `json:"reference"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", fmt.Errorf("failed to decode SantimPay response: %v", err)
	}

	return response.PaymentURL, response.Reference, nil
}