package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
)

func generateSecretHash(username string) string {
	digest := ""
	clientID := os.Getenv("AWS_COGNITO_CLIENT_ID")
	clientSecret := os.Getenv("AWS_COGNITO_CLIENT_SECRET")

	if clientID != "" && clientSecret != "" {
		mac := hmac.New(sha256.New, []byte(clientSecret))
		mac.Write([]byte(username + clientID))

		digest = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	}

	return digest
}
