package gateway

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateSignedToken(amount float64, reason, merchantID string) (string, error) {
	privateKeyPEM := os.Getenv("SANTIMPAY_PRIVATE_KEY")

	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to parse private key PEM")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse ec private key: %v", err)
	}

	claims := jwt.MapClaims{
		"amount":        amount,
		"paymentReason": reason,
		"merchantId":    merchantID,
		"generated":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privateKey)
}
